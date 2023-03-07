package global

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient  *mongo.Client
	PostgresqlDB *sql.DB
	RedisDB      *redis.Client
	Trans        ut.Translator
)
