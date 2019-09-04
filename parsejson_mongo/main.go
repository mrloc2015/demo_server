package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type People struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Age       int    `json:"age" bson:"age"`
	Married   bool   `json:"married" bson:"married"`
}

func main() {

	url := "http://127.0.0.1:9090/v1/people"

	res, err := http.Get(url)

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var peopleListing []People
	err = decoder.Decode(&peopleListing)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
		case *json.SyntaxError:
			fmt.Println("Json syntax Error: " + v.Error())
		}
	} else {
		fmt.Println("===========Result From People API===============")
		for i, info := range peopleListing {
			//  @TODO print all peopleListing from json
			fmt.Printf("%d: %s %s %d %v\n", i, info.FirstName, info.LastName, info.Age, info.Married)
		}

		//  @TODO Writing peopleListing to mongoDB
		clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MongoDB!")

		peopleCollection := client.Database("practice1").Collection("people")

		var prepareData []interface{}
		for _, v := range peopleListing{
			prepareData = append(prepareData, v)
		}

		insertManyResult, err := peopleCollection.InsertMany(context.TODO(), prepareData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

		//fmt.Println(peopleListing)
	}

}
