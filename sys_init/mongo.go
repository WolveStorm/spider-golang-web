package sys_init

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"spider-golang-web/global"
	"time"
)

func InitMongo() {
	var err error
	//1.建立连接
	if global.MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", global.Host)).SetConnectTimeout(5*time.Second)); err != nil {
		zap.S().Errorf("error to start mongo,detail:%v", err)
		return
	}
}
