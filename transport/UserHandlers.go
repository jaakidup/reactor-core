package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jaakidup/reactor-core/model"
)

func getUserByID(w http.ResponseWriter, r *http.Request) {

	ID := chi.URLParam(r, "id")
	user, err := Service.User.GetUser(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := Service.User.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := Service.User.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := struct {
		ID   string
		Note string
	}{ID: id, Note: "User created"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
