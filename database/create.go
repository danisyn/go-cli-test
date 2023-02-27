package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"web/slack"
)

type Users struct {
	User  string
	Password string
	Registered string
	Token string
}

func CreateUser(user string, password string, base64File string) bool{
	var status bool
	uri := "mongodb://" + username + ":" + pass + "@" + host +":" + port + "/app"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
			panic(err)
	}
	usersCollection := client.Database("app").Collection("users")
	timestamp := time.Now()
	registered := fmt.Sprint(timestamp)
	token := GenerateToken()
	insert := Users{User: user, Password: password, Registered: registered, Token: token}

	_ , err = usersCollection.InsertOne(context.TODO(), insert)
	if err != nil {
		status = false
	} else {
		status = true
		GenerateFile(user, token, base64File)
		fmt.Printf("Save this token: %s\n",token)
		slack.PostToSlack(user, token, base64File)
	}
	return status
}