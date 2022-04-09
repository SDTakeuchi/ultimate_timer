package config


/*

WIP
postgres to mongo

*/


// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// func NewMongoDB() *mongo.Database {
// 	// コンテキストの作成
// 	//   - バックグラウンドで接続する。タイムアウトは10秒
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	// 関数を抜けたらクローズするようにする
// 	defer cancel()
// 	// 指定したURIに接続する
// 	c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	defer c.Disconnect(ctx)
// 	// DBにPingする
// 	err = c.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		fmt.Println("connection error:", err)
// 	} else {
// 		fmt.Println("connection success")
// 	}
// 	return c
// }

// type MongoConfig struct {
// 	DBname   string
// 	Username string
// 	Password string
// 	Host     string
// 	Port     int
// 	Timeout  int
// }

// var mongoConfig MongoConfig = MongoConfig{
// 	DBname:   "mongodb",
// 	Username: "root",
// 	Password: "password",
// 	Host:     "mongo",
// 	Port:     27017,
// 	Timeout:  10,
// }

// func Connect() (*mongo.Database, error) {
// 	connPattern := "mongodb://%v:%v@%v:%v"
// 	if mongoConfig.Username == "" {
// 		connPattern = "mongodb://%s%s%v:%v"
// 	}

// 	clientUrl := fmt.Sprintf(connPattern,
// 		mongoConfig.Username,
// 		mongoConfig.Password,
// 		mongoConfig.Host,
// 		mongoConfig.Port,
// 	)
// 	clientOptions := options.Client().ApplyURI(clientUrl)
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), time.Duration(mongoConfig.Timeout)*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.Database(mongoConfig.DBname), err
// }
