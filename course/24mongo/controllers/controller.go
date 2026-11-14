package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "moon"

const db = "netflix"
const colName = "watchList"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client , err := mongo.Connect(context.TODO(), clientOption)

	if err!= nil{
		log.Panic(err)
	}
	fmt.Println("mongo connected")

	collection =  client.Database(db).Collection(colName)
	fmt.Println("Collection is ready")
}
