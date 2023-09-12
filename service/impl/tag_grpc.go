package impl

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/service"
	"github.com/wutipong/mangaweb3-backend/tag"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type TagServer struct {
	service.UnimplementedTagServiceServer
	EntClient *ent.Client
}

const (
	DefaultTagPerPage = 30
)

func (s *TagServer) List(ctx context.Context, req *service.TagListRequest) (resp *service.TagListResponse, err error) {
	log.Info().Interface("request", req).Msg("Tag list")

	itemPerPage := req.ItemPerPage
	if itemPerPage == 0 {
		itemPerPage = DefaultTagPerPage
	}

	allTags, err := tag.ReadPage(ctx, req.FavoriteOnly, int(req.Page), int(itemPerPage))
	if err != nil {
		return
	}

	total, err := tag.Count(ctx, req.FavoriteOnly)
	if err != nil {
		return
	}

	resp = &service.TagListResponse{
		TotalPage: (int32(total) / itemPerPage) + 1,
	}

	for _, t := range allTags {
		resp.Tags = append(resp.Tags, &service.Tag{
			Id:        int64(t.ID),
			Name:      t.Name,
			Favorite:  t.Favorite,
			Thumbnail: wrapperspb.Bytes(t.Thumbnail),
		})
	}

	return
}
