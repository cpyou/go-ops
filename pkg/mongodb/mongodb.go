package mongodb

import (
	"context"
	"fmt"
	"go-ops/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoClient struct {
	Client *mongo.Client
	Dbname string
}


func GetMongoClient() *mongo.Client {
	c := config.GetConfig().MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(c.Uri))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return client
}

func (m *MongoClient) InsertOne(c string, data interface{}) (*mongo.InsertOneResult, error) {
	collection := m.Client.Database(m.Dbname).Collection(c)
	insertOneResult, err := collection.InsertOne(context.TODO(), data)
	return insertOneResult, err
}

func (m *MongoClient) InsertMany(c string, data []interface{}) (*mongo.InsertManyResult, error) {
	collection := m.Client.Database(m.Dbname).Collection(c)
	insertManyResult, err := collection.InsertMany(context.TODO(), data)
	return insertManyResult, err
}

func (m *MongoClient) UpdateOne(c string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := m.Client.Database(m.Dbname).Collection(c)
	updateOneResult, err := collection.UpdateOne(context.TODO(), filter, update)
	return updateOneResult, err
}

func (m *MongoClient) UpdateMany(c string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := m.Client.Database(m.Dbname).Collection(c)
	updateOneResult, err := collection.UpdateMany(context.TODO(), filter, update)
	return updateOneResult, err
}

func (m *MongoClient) FindOne(c string, filter interface{}, result interface{}) error {
	collection := m.Client.Database(m.Dbname).Collection(c)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return err
}

func (m *MongoClient) Find(c string, filter interface{}) (bson.M, error) {
	collection := m.Client.Database(m.Dbname).Collection(c)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(context.TODO(), filter)
	defer cur.Close(ctx)
	result := bson.M{}
	for cur.Next(ctx) {
		var temp bson.M
		err := cur.Decode(&temp)
		if err != nil {log.Fatal(err)}
		for k, v := range(temp) {
			result[k] = v
		}
	}
	return result, err
}