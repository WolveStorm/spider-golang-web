package sys_init

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"spider-golang-web/global"
)

func InitPostgres() {
	var err error
	dsn := fmt.Sprintf("host=%s port=5432 user=postgres password=123456 sslmode=disable database=game_info", global.Host)
	if global.PostgresqlDB, err = sql.Open("postgres", dsn); err != nil {
		zap.S().Errorf("error to start postgresql,detail:%v", err)
		return
	}
	fmt.Println(global.PostgresqlDB.Ping())
}
