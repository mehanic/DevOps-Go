package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kirigaikabuto/Golang1300Lessons/FAQbot/bot"
	"log"
	"net/http"
	"os"
)

var (
	mongoDb         = ""
	mongoUrl        = ""
	mongoCollection = ""
)

func parseEnv() {
	err := godotenv.Overload("config.env")
	if err != nil {
		log.Fatal(err)
		return
	}
	mongoDb = os.Getenv("MONGO_DB")
	mongoUrl = os.Getenv("MONGO_URL")
	mongoCollection = os.Getenv("MONGO_COLLECTION")
}

func main() {
	parseEnv()
	config := bot.MongoConfig{
		Url:        mongoUrl,
		Database:   mongoDb,
		Collection: mongoCollection,
	}
	botMongoStore, err := bot.NewBotMongoStore(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	botService := bot.NewBotService(botMongoStore)
	botHttpEndoints := bot.NewHttpBotEndpoints(botService)
	r := mux.NewRouter()
	r.Methods("POST").Path("/bot").HandlerFunc(botHttpEndoints.MakeCreateBot())
	r.Methods("GET").Path("/bot/{id}").HandlerFunc(botHttpEndoints.MakeGetBotById("id"))
	r.Methods("GET").Path("/bot").HandlerFunc(botHttpEndoints.MakeListBot())
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", r)
	//bots, err := botMongoStore.List()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(bots)
	//b1 := &bot.Bot{
	//	Id:       "",
	//	Name:     "asdsadd",
	//	Telegram: "asdadasd",
	//	Facebook: "asdadad",
	//	WhatsApp: "asdsadad",
	//}
	//res, err := botMongoStore.Create(b1)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(res)
}
