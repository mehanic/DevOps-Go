package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

// struct for storing data
type task struct {
	ID          string `json:id bson:_id`
	Description string `json:description bson:description`
	Category    string `json:category bson:category`
	DueDate     time.Time `json:dueDate bson:dueDate`
}

const (
	layoutISO = "2006-01-02"
)

var taskCollection = db().Database("db_task").Collection("tasks")

// Create task
func createTask(w http.ResponseWriter, r *http.Request) {
	type createdBody struct {
		Description string `json:description`
		Category    string `json:category`
		DueDate     string `json:dueDate`
	}
	var body createdBody
	var task1 task
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if len(body.Description) == 0 {
		respondWithError(w, http.StatusBadRequest, "The field description is required.")
		return
	}
	if len(body.Category) == 0 {
		respondWithError(w, http.StatusBadRequest, "The field category is required.")
		return
	}
	if len(body.DueDate) == 0 {
		respondWithError(w, http.StatusBadRequest, "The field due date is required.")
		return
	} else {
		date, er := time.Parse(layoutISO, body.DueDate)
		if er != nil {
			respondWithError(w, http.StatusBadRequest, er.Error())
			return
		}
		now := time.Now()
		fmt.Println(date)
		fmt.Println(now)
		if now.After(date) {
			respondWithError(w, http.StatusBadRequest, "The field due date is smaller than " + now.Format(layoutISO))
			return
		}
		task1.DueDate = date
	}
	task1.Description = body.Description
	task1.Category = body.Category
	insertResult, err := taskCollection.InsertOne(context.TODO(), task1)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusCreated, insertResult)
}

//List tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	var results []primitive.M
	cur, err := taskCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request" + err.Error())
		return
	}
	for cur.Next(context.TODO()) {
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		results = append(results, elem)
	}
	cur.Close(context.TODO())
	resultJson := map[string][]primitive.M{
		"result": results,
	}
	respondWithJSON(w, http.StatusOK, resultJson)
}

//Update Task
func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request " + err.Error())
		return
	}
	type updateBody struct {
		Description string `json:description`
		Category    string `json:category`
		DueDate     string `json:dueDate`
	}
	var body updateBody
	var task1 task
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if len(body.Description) == 0 {
		respondWithError(w, http.StatusBadRequest, "The field description is required.")
		return
	}
	if len(body.Category) == 0 {
		respondWithError(w, http.StatusBadRequest, "The field category is required.")
		return
	}
	if len(body.DueDate) == 0 {
		respondWithError(w, http.StatusBadRequest, "The field due date is required.")
		return
	} else {
		date, er := time.Parse(layoutISO, body.DueDate)
		if er != nil {
			respondWithError(w, http.StatusBadRequest, er.Error())
			return
		}
		now := time.Now()
		if date.Before(now) {
			respondWithError(w, http.StatusBadRequest, "The field due date is smaller than " + now.Format(layoutISO))
			return
		}
		task1.DueDate = date
	}
	task1.Description = body.Description
	task1.Category = body.Category
	filter := bson.D{{"_id", _id}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{
		{"$set", bson.D{{"description", task1.Description}}},
		{"$set", bson.D{{"category", task1.Category}}},
		{"$set", bson.D{{"duedate", task1.DueDate}}},
	}
	updateResult := taskCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)
	var result primitive.M
	_ = updateResult.Decode(&result)
	if result == nil {
		resultJson := map[string]primitive.M{
			"result": result,
		}
		respondWithJSON(w, http.StatusOK, resultJson)
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}
}

//Delete Task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request " + err.Error())
		return
	}
	opts := options.Delete().SetCollation(&options.Collation{})
	res, err1 := taskCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err1 != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resultJson := map[string]string{
		"result": "Deleted " + strconv.Itoa(int(res.DeletedCount)) + " documents",
	}
	respondWithJSON(w, http.StatusOK, resultJson)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
