package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var host = "localhost"
var port = "27017"
var username = "develop"
var pass = "develop123"

func PingDb() bool{
	var status bool
	uri := "mongodb://" + username + ":" + pass + "@" + host +":" + port + "/app"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
			panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
        fmt.Println("💀 Host unreachable 💀")
		status = false
	}else {
		status = true
	}
	return status
}