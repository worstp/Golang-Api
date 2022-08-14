package userhandlers

import (
	"encoding/json"
	"fmt"
	"go/api/models/user"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var model user.User

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := user.PostUserService(model)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error on insert user: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("User inserted successfully, ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func List(w http.ResponseWriter, r *http.Request) {
	users, err := user.GetAllUsersService(0)
	if err != nil {
		log.Printf("Error on get all users: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}