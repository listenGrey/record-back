package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	RecordCol *mongo.Collection
)

func init() {
	connCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // 建议确保资源及时释放

	// 替换为你自己的 MongoDB 连接 URI
	client, err := mongo.Connect(connCtx, options.Client().
		ApplyURI("mongodb://admin:tinghui0430@localhost:27017/?authSource=admin"))
	// ApplyURI("mongodb://mongo:mlybQjXAeORwDwuzWJHaLTgjEPEnqVyH@metro.proxy.rlwy.net:59428"))
	if err != nil {
		log.Fatal("连接 MongoDB 失败:", err)
	}

	// 选择数据库和集合
	RecordCol = client.Database("record").Collection("trade")
}
