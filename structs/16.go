package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func InsertData(collection *mongo.Collection, product Product) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ListData(collection *mongo.Collection) ([]Product, error) {
	filter := bson.D{}
	products := []Product{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetDataByName(collection *mongo.Collection, name string) (*Product, error) {
	filter := bson.D{{"name", name}}
	p := &Product{}
	err := collection.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func UpdateData(collection *mongo.Collection, name string, product Product) error {
	filter := bson.D{{"name", name}}
	//collection.DeleteOne(context.TODO(),filter)
	update := bson.D{
		{
			"$set",
			bson.D{
				{"name", product.Name},
				{"price", product.Price},
			},
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	url := "mongodb://localhost:27017"
	databaseName := "db11"
	collectionName := "collection11"
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	db := client.Database(databaseName)
	collection := db.Collection(collectionName)
	//p1 := Product{
	//	Name:  "132213",
	//	Price: 2323,
	//}
	//_, err = InsertData(collection, p1)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//products, err := ListData(collection)
	//for i, v := range products {
	//	fmt.Println(i, v.Name)
	//}
	//pr, err := GetDataByName(collection, "132213")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(pr)
	//err = UpdateData(collection, "11111", Product{
	//	Name:  "222222",
	//	Price: 2222,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

}
