package bot

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type HttpBotEndpoints interface {
	MakeCreateBot() func(w http.ResponseWriter, r *http.Request)
	MakeGetBotById(idParam string) func(w http.ResponseWriter, r *http.Request)
	MakeListBot() func(w http.ResponseWriter, r *http.Request)
}

type httpBotEndpoints struct {
	service BotService
}

func NewHttpBotEndpoints(s BotService) HttpBotEndpoints {
	return &httpBotEndpoints{service: s}
}

func (h *httpBotEndpoints) MakeCreateBot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := &CreateBotCommand{}
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
		bot, err := h.service.CreateBot(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusCreated, bot)
		return
	}
}

func (h *httpBotEndpoints) MakeGetBotById(idParam string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := mux.Vars(r)
		id, ok := data[idParam]
		if !ok {
			respondJSON(w, http.StatusBadRequest, ErrorMessage{
				Message: "Please send id",
				Status:  http.StatusBadRequest,
			})
		}
		cmd := &GetBotByIdCommand{Id: id}
		bot, err := h.service.GetBotById(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusOK, bot)
		return
	}
}

func (h *httpBotEndpoints) MakeListBot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := &ListBotCommand{}
		bots, err := h.service.ListBot(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusOK, bots)
		return
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
