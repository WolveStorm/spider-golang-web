package proto

import (
	"context"
)

type ServiceGameImpl struct{}

func (s *ServiceGameImpl) GameList(ctx context.Context, req *GameListFilterRequest) (*GameListResponse, error) {
	// 尝试从缓存中返回信息
	resp, ok := FindGameListInRedis(req.Page, req.PageSize, req.Keyword)
	if ok {
		return resp, nil
	}
	// 缓存失效,重新更新
	GameListToRedis()
	// 如果没有从redis获得，则走数据库
	var (
		page     int32 = 1
		pageSize int32 = 10
	)
	if req.Page != 0 {
		page = req.Page
	}
	if req.PageSize != 0 {
		pageSize = req.PageSize
	}
	offset := (page - 1) * pageSize
	gameList, total, err := FindGameList(offset, pageSize, req.Keyword)
	if err != nil {
		return nil, err
	}
	return &GameListResponse{
		Total: total,
		List:  gameList,
	}, nil
}
func (s *ServiceGameImpl) GameDetail(ctx context.Context, req *GameDetailRequest) (*GameDetailInfoResp, error) {
	// 尝试从缓存中返回信息
	resp, ok := FindGameDetailInRedis(req.GameName)
	if ok {
		return resp, nil
	}
	// 缓存失效,重新更新
	GameDetailToRedis(req.GameName)
	detail, err := FindGameDetail(req.GameName)
	if err != nil {
		return nil, err
	}
	return detail, nil
}
func (s *ServiceGameImpl) mustEmbedUnimplementedGameServer() {}
