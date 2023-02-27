package database

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

type ConfFile struct {
	Token string
	Registered string
	File string
}

func GenerateFile(user string, token string, base64File string) bool{
	var status bool
	uri := "mongodb://" + username + ":" + pass + "@" + host +":" + port + "/app"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
			panic(err)
	}
	usersCollection := client.Database("app").Collection("kubeconfigs")
	timestamp := time.Now()
	registered := fmt.Sprint(timestamp)

	insert := ConfFile{Token: token, Registered: registered, File: base64File}

	_ , err = usersCollection.InsertOne(context.TODO(), insert)
	if err != nil {
		status = false
	} else {
		status = true
	}
	return status
}