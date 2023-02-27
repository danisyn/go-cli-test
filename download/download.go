package download

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"web/login"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Download struct {
	Token string
}

type File struct {
	File string
}

var host = "localhost"
var port = "27017"
var username = "develop"
var pass = "develop123"

func DownloadConfFile() bool{
	var status bool
	exists, token := login.LogIn()
	if exists == true {
		
		uri := "mongodb://" + username + ":" + pass + "@" + host +":" + port + "/app"
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
				panic(err)
		}

		usersCollection := client.Database("app").Collection("kubeconfigs")
		filter := Download{Token: token}
		var result File

		err = usersCollection.FindOne(context.TODO(), filter).Decode(&result)
		if result.File == "" {
			status = false
			fmt.Println("There is no Kubeconfig file configured for this user")
		} else {
			status = true
			HOME := os.Getenv("HOME")
			file := HOME + "/.kube/config"
			rawDecodedText, _ := base64.StdEncoding.DecodeString(result.File)
			f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
			panic(err) // i'm simplifying it here. you can do whatever you want.
			}
			f.WriteString(string(rawDecodedText))
			fmt.Printf("File saved in %s/.kube/config\n", HOME)
		}
		
	}
	
	
return status


}