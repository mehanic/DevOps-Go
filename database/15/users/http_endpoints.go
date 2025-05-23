package users

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/Golang1300Lessons/15/redis_connect"
	"io/ioutil"
	"net/http"
	"time"
)

type UsersHttpEndpoints interface {
	Login() func(w http.ResponseWriter, r *http.Request)
	GetInfo() func(w http.ResponseWriter, r *http.Request)
	Register() func(w http.ResponseWriter, r *http.Request)
}

type usersHttpEndpoints struct {
	//mongoStore
	usersStore UsersStore
	//redisStore
	redisConnect *redis_connect.RedisConnectStore
}

func NewUsersHttpEndpoints(u UsersStore, r *redis_connect.RedisConnectStore) UsersHttpEndpoints {
	return &usersHttpEndpoints{
		usersStore:   u,
		redisConnect: r,
	}
}

func (u *usersHttpEndpoints) Login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//getting data from request body
		cmd := &LoginRequestCommand{}
		dataJson, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		err = json.Unmarshal(dataJson, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		//service.go
		//getting data about user from db
		user, err := u.usersStore.GetByUsernameAndPassword(cmd.Username, cmd.Password)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		//creating token
		token := uuid.New().String()
		expirationTime := 5 * time.Minute
		err = u.redisConnect.Save(token, user.Id, expirationTime)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		response := &LoginResponse{
			AccessToken:    token,
			ExpirationTime: expirationTime.String(),
		}
		respondJSON(w, http.StatusOK, response)
	}
}

func (u *usersHttpEndpoints) Register() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//getting data from request body
		cmd := &LoginRequestCommand{}
		dataJson, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		err = json.Unmarshal(dataJson, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		//service.go
		user, err := u.usersStore.Create(&User{
			Username: cmd.Username,
			Password: cmd.Password,
		})
		respondJSON(w, http.StatusCreated, user)
	}
}

func (u *usersHttpEndpoints) GetInfo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		contextUserId := r.Context().Value("user_id")
		userId := contextUserId.(string)
		user, err := u.usersStore.Get(userId)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusOK, user)
	}
}

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
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
