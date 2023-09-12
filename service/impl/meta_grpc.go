package impl

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
	ent_meta "github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/service"
	"github.com/wutipong/mangaweb3-backend/tag"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type MetaServer struct {
	service.UnimplementedMetaServiceServer
	EntClient *ent.Client
}

const (
	DefaultItemPerPage = 30
)

func (s *MetaServer) List(ctx context.Context, req *service.MetaListRequest) (resp *service.MetaListResponse, err error) {
	var sortBy meta.SortField
	switch req.Sort {
	case service.SORT_BY_SORT_BY_NAME:
		sortBy = meta.SortFieldName

	case service.SORT_BY_SORT_BY_CREATE_TIME:
		sortBy = meta.SortFieldCreateTime
	}

	var order meta.SortOrder
	switch req.Order {
	case service.ORDER_ORDER_ASCENDING:
		order = meta.SortOrderAscending

	case service.ORDER_ORDER_DESCENDING:
		order = meta.SortOrderDescending
	}

	allMeta, err := meta.SearchItems(
		ctx,
		req.Search,
		req.FavoriteOnly,
		req.Tag,
		sortBy,
		order,
		int(req.Page),
		int(req.ItemPerPage))
	if err != nil {
		return
	}

	count, err := meta.CountItems(ctx,
		req.Search,
		req.FavoriteOnly,
		req.Tag,
		sortBy,
		order)

	if err != nil {
		return
	}

	itemPerPage := req.ItemPerPage
	if itemPerPage == 0 {
		itemPerPage = DefaultItemPerPage
	}
	pageCount := int32(count) / req.ItemPerPage
	if int32(count)%req.ItemPerPage > 0 {
		pageCount++
	}

	if req.Page > pageCount || req.Page < 0 {
		req.Page = 0
	}

	log.Info().
		Interface("request", req).
		Msg("Browse")

	resp = &service.MetaListResponse{
		TotalPage: pageCount,
	}
	resp.Items = make([]*service.Meta, len(allMeta))

	for i := range resp.Items {
		resp.Items[i] = &service.Meta{
			Id:         int64(allMeta[i].ID),
			Name:       allMeta[i].Name,
			CreateTime: timestamppb.New(allMeta[i].CreateTime),
			Favorite:   allMeta[i].Favorite,
			Thumbnail:  wrapperspb.Bytes(allMeta[i].Thumbnail),
			Read:       allMeta[i].Read,
		}
	}

	if req.Tag != "" {
		tagObj, e := tag.Read(ctx, req.Tag)
		if e != nil {
			err = e
			return
		}

		resp.TagFavorite = tagObj.Favorite
	}

	return
}

func (s *MetaServer) Get(ctx context.Context, req *service.MetaGetRequest) (resp *service.MetaGetResponse, err error) {
	m, err := s.EntClient.Meta.Get(ctx, int(req.Id))
	if err != nil {
		return
	}

	if !m.Read {
		m.Read = true
		meta.Write(ctx, m)
	}

	log.Info().
		Interface("request", req).
		Msg("View Item")

	tags, err := m.QueryTags().All(ctx)
	if err != nil {
		return
	}

	resp = &service.MetaGetResponse{
		Item: &service.Meta{
			Id:         int64(m.ID),
			Name:       m.Name,
			CreateTime: timestamppb.New(m.CreateTime),
			Favorite:   m.Favorite,
			Thumbnail:  wrapperspb.Bytes(m.Thumbnail),
			Read:       m.Read,
		},
		PageCount: int32(len(m.FileIndices)),
	}

	for _, t := range tags {
		resp.Tags = append(resp.Tags, &service.Tag{
			Name:     t.Name,
			Favorite: t.Favorite,
		})
	}

	return
}

func (s *MetaServer) SetFavorite(ctx context.Context, req *service.SetFavoriteRequest) (resp *service.SetFavoriteResponse, err error) {
	err = s.EntClient.Meta.Update().
		Where(ent_meta.ID(int(req.Id))).
		SetFavorite(req.Favorite).
		Exec(ctx)

	if err != nil {
		return
	}

	resp = &service.SetFavoriteResponse{
		Favorite: req.Favorite,
	}

	return
}

func (s *MetaServer) GetPage(ctx context.Context, req *service.GetPageRequest) (resp *service.GetPageResponse, err error) {
	m, err := s.EntClient.Meta.Get(ctx, int(req.Id))
	if err != nil {
		return
	}

	data, f, err := OpenZipEntry(m, int(req.Page))
	if err != nil {
		return
	}

	resp = &service.GetPageResponse{
		ImageData: wrapperspb.Bytes(data),
	}

	switch filepath.Ext(strings.ToLower(f)) {
	case ".jpg", ".jpeg":
		resp.ContentType = "image/jpeg"
	case ".png":
		resp.ContentType = "image/png"
	case ".webp":
		resp.ContentType = "image/webp"
	}

	return
}

func OpenZipEntry(m *ent.Meta, index int) (content []byte, filename string, err error) {
	if len(m.FileIndices) == 0 {
		err = fmt.Errorf("image file not found")
	}

	fullpath := filepath.Join(meta.BaseDirectory, m.Name)
	r, err := zip.OpenReader(fullpath)
	if err != nil {
		return
	}

	defer r.Close()

	zf := r.File[m.FileIndices[index]]

	if zf == nil {
		err = fmt.Errorf("file not found : %v", index)
		return
	}

	filename = zf.Name
	reader, err := zf.Open()
	if err != nil {
		return
	}
	defer reader.Close()
	if content, err = io.ReadAll(reader); err != nil {
		content = nil
		return
	}
	return
}
