package main

import (
	"log"
	"github.com/redistest/server"
	"github.com/redistest/app"
	_ "github.com/lib/pq"
)

func main() {
	db, err := app.DBConnectionBuilder()
	redis := app.RedisConnectionBuilders()
	
	if err != nil {
		log.Panicf("Some error with connetion to DB: %s", err.Error())
	}

	handler := app.HandlerBuilder(db, redis)
	router := handler.RouterBuilder()

	err = server.SeverConnectionBuilder(router)
	if err != nil {
		log.Panicf("Some error with Run to DB: %s", err.Error())
	}
}