package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Filter struct {
	User string
	Password string
	Token string
}

func GetUser(user string, password string, token string) bool{
	var status bool
	uri := "mongodb://" + username + ":" + pass + "@" + host +":" + port + "/app"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
			panic(err)
	}
	usersCollection := client.Database("app").Collection("users")
	filter := Filter{User:user, Password: password, Token: token}
	var result Filter
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if result.User == "" {
		status = false
	} else {
		status = true
	}

	return status
}