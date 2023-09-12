package impl

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
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

func (s *MetaServer) List(ctx context.Context, req *service.ListRequest) (resp *service.ListResponse, err error) {
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

	resp = &service.ListResponse{
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

func (s *MetaServer) Get(ctx context.Context, req *service.GetRequest) (resp *service.GetResponse, err error) {
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

	resp = &service.GetResponse{
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
