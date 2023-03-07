package proto

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"spider-golang-web/global"
	"strings"
	"time"
)

func GameListToRedis() {
	// 查询全部列表
	gameList, _, err := FindGameList(0, 0, "")
	if err != nil {
		return
	}
	jInfo, err := json.Marshal(gameList)
	if err != nil {
		zap.S().Error("error to marshal list cache")
		return
	}
	// 重新设置缓存，并且这一次的查询从数据库查找
	_, err = global.RedisDB.SetEx(context.Background(), global.KVGameList, string(jInfo), 60*time.Minute*24*7).Result()
	if err != nil {
		zap.S().Error("set list cache error")
		return
	}
}

func FindGameListInRedis(reqPage, reqPageSize int32, reqKeyword string) (*GameListResponse, bool) {
	// 直接在缓存中查询
	cacheList := make([]*GameDetailInfoResp, 0)
	jsonListStr, err := global.RedisDB.Get(context.Background(), global.KVGameList).Result()
	if err != nil {
		zap.S().Errorf("error to read cache,detail:%v", err)
		return nil, false
	}
	err = json.Unmarshal([]byte(jsonListStr), &cacheList)
	if err != nil {
		zap.S().Error("list json unmarshall to struct error")
		return nil, false
	}
	// 处理传入参数
	var (
		page     int32 = 1
		pageSize int32 = 10
		keyword        = ""
	)
	if reqPage != 0 {
		page = reqPage
	}
	if reqPageSize != 0 {
		pageSize = reqPageSize
	}
	if reqKeyword != "" {
		keyword = reqKeyword
	}
	// 查找到了cacheList，直接从redis读取数据，否则仍然走数据库
	filterList := make([]*GameDetailInfoResp, 0)
	// 基于列表找到合适的返回
	for _, item := range cacheList {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		filterList = append(filterList, item)
	}
	total := int32(len(filterList))
	resp := &GameListResponse{
		Total: total,
	}
	offset := (page - 1) * pageSize
	if offset >= int32(len(cacheList)) {
		resp.List = make([]*GameDetailInfoResp, 0)
		return resp, true
	}
	end := offset + pageSize
	if end >= int32(len(cacheList)) {
		end = int32(len(cacheList))
	}
	resp.List = filterList[offset:end]
	return resp, true
}

func GameDetailToRedis(gameName string) {
	// 查询全部列表
	detail, err := FindGameDetail(gameName)
	if err != nil {
		return
	}
	jInfo, err := json.Marshal(detail)
	if err != nil {
		zap.S().Error("error to marshal game detail cache")
		return
	}
	// 重新设置缓存，并且这一次的查询从数据库查找
	_, err = global.RedisDB.Set(context.Background(), global.KVGameDetail+gameName, string(jInfo), 60*time.Minute*24*7).Result()
	if err != nil {
		zap.S().Error("set game detail cache error")
		return
	}
}
func FindGameDetailInRedis(gameName string) (*GameDetailInfoResp, bool) {
	detail := &GameDetailInfoResp{}
	detailJsonStr, err := global.RedisDB.Get(context.Background(), global.KVGameDetail+gameName).Result()
	if err != nil {
		zap.S().Error("error to read cache")
		return nil, false
	}
	err = json.Unmarshal([]byte(detailJsonStr), detail)
	if err != nil {
		zap.S().Error("list json unmarshall to struct error")
		return nil, false
	}
	return detail, true
}
