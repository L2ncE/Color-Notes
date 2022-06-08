package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"wechat/global"
)

var mongoDB *mongo.Client

func InitMongoDB() error {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", global.Settings.MongoDBInfo.Host, global.Settings.MongoDBInfo.Port))
	// 连接 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	mongoDB = client
	return nil
}
