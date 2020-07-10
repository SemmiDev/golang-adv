package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database("belajar_golang"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Sam", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Dev", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func find() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Sam"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]student, 0)
	for csr.Next(ctx) {
		var row student
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name  :", result[0].Name)
		fmt.Println("Grade :", result[0].Grade)
	}
}

func update() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "Dev"}
	var changes = student{"Adisti", 2}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}

func delete() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.error())
	}

	var selector = bson.M{"name:" "Dev"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
		if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")
}

func main() {
	insert()
	find()
	update()
	delete()

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Total :", result[0]["Total"])
	}

}
