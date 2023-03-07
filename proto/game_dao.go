package proto

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"spider-golang-web/global"
)

func FindGameList(offset int32, pageSize int32, keyword string) ([]*GameDetailInfoResp, int32, error) {
	sql := "select name,avatar_url,company,score,download_times,apk_url,description from gameinfo"
	if keyword != "" {
		sql += " where name like '%" + keyword + "%'"
	}
	rows, err := global.PostgresqlDB.Query(sql)
	if err != nil {
		return nil, 0, status.Errorf(codes.Internal, "service unavailable")
	}
	list := make([]*GameDetailInfoResp, 0)
	for rows.Next() {
		info := &GameDetailInfoResp{}
		err = rows.Scan(&info.Name, &info.AvatarUrl, &info.Company, &info.Score, &info.DownloadTimes, &info.ApkUrl, &info.Desc)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, info)
	}
	// 如果没有传入pagesize和offset，则查询全部列表的功能
	if offset == 0 && pageSize == 0 {
		return list, int32(len(list)), err
	}
	if offset >= int32(len(list)) {
		return make([]*GameDetailInfoResp, 0), int32(len(list)), nil
	}
	end := offset + pageSize
	if end >= int32(len(list)) {
		end = int32(len(list))
	}
	return list[offset:end], int32(len(list)), nil
}

func FindGameDetail(gameName string) (*GameDetailInfoResp, error) {
	sql := "select name,avatar_url,company,score,download_times,apk_url,description from gameinfo"
	if gameName != "" {
		sql += " where name = '" + gameName + "'"
	} else {
		return nil, errors.New("you should give arg like game_name let me know which game you want to query")
	}
	row := global.PostgresqlDB.QueryRow(sql)
	info := &GameDetailInfoResp{}
	err := row.Scan(&info.Name, &info.AvatarUrl, &info.Company, &info.Score, &info.DownloadTimes, &info.ApkUrl, &info.Desc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "service unavailable")
	}
	return info, nil
}
