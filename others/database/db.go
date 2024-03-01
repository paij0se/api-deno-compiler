package database

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const alphanumericChars = "abcdefghijklmnopqrstuvwxyz0123456789"

func generateRandomString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = alphanumericChars[rand.Intn(len(alphanumericChars))]
	}
	return string(result)
}

func DbInsert(client *mongo.Client, code string) string {
	collection := client.Database("adc").Collection("codes")

	id := generateRandomString(6)
	_, err := collection.InsertOne(context.Background(), bson.M{"_id": id, "code": code})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Code inserted successfully With ID:", id)

	return id
}

func DbGet(client *mongo.Client, id string) string {
	collection := client.Database("adc").Collection("codes")

	var result bson.M
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return "Code Ain't Found"
	}
	fmt.Println("Code retrieved successfully")

	code := result["code"].(string)
	return code
}

func ConnectToDB() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DATABASE_URL")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
