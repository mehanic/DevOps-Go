package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type MegaHelloResponse struct {
	Message string `json:"message"`
}

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type HelloRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func CreateProduct(data *HelloRequest) *Product {
	return &Product{
		Id:    "1",
		Name:  data.Name,
		Price: data.Price,
	}
}

func MegaHelloHandler(param1, param2 string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//read params from path /hello/{config}
		data := mux.Vars(r)
		fmt.Println(data)
		config, ok := data[param1]
		if !ok {
			respondJSON(w, http.StatusBadRequest, &ErrorMessage{
				Message: "need parameter1",
				Status:  http.StatusBadRequest,
			})
			return
		}
		value, ok := data[param2]
		if !ok {
			respondJSON(w, http.StatusBadRequest, &ErrorMessage{
				Message: "need parameter2",
				Status:  http.StatusBadRequest,
			})
			return
		}
		message := config + value
		a := r.URL.Query().Get("a")
		if a != "" {
			message += a
		}
		respondJSON(w, 200, &MegaHelloResponse{Message: message})
	}
}

func HelloHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, &ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
			return
		}
		req := &HelloRequest{}
		err = json.Unmarshal(data, &req)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, &ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
			return
		}
		product := CreateProduct(req)
		respondJSON(w, 200, product)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/hello/{config}/{value}").HandlerFunc(MegaHelloHandler("config", "value"))
	r.Methods("POST").Path("/hello").HandlerFunc(HelloHandler())
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
