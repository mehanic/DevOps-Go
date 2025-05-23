package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	//"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter6/mongodb"
)

func main() {
	if err := Exec("mongodb://localhost"); err != nil {
		panic(err)
	}
}

// Setup initializes a mongo client
func Setup(ctx context.Context, address string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// cancel will be called when setup exits
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	return client, nil
}

//--

// State is our data model
type State struct {
	Name       string `bson:"name"`
	Population int    `bson:"pop"`
}

// Exec creates then queries an Example
func Exec(address string) error {
	ctx := context.Background()
	db, err := Setup(ctx, address)
	if err != nil {
		return err
	}

	coll := db.Database("gocookbook").Collection("example")

	vals := []interface{}{&State{"Washington", 7062000}, &State{"Oregon", 3970000}}

	// we can inserts many rows at once
	if _, err := coll.InsertMany(ctx, vals); err != nil {
		return err
	}

	var s State
	if err := coll.FindOne(ctx, bson.M{"name": "Washington"}).Decode(&s); err != nil {
		return err
	}

	if err := coll.Drop(ctx); err != nil {
		return err
	}

	fmt.Printf("State: %#v\n", s)
	return nil
}
