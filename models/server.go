package models

import (
	fileadapter "github.com/casbin/casbin/persist/file-adapter"
	"github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	DB          *mongo.Client
	RedisCli    *redis.Client
	FileAdapter *fileadapter.Adapter
}
