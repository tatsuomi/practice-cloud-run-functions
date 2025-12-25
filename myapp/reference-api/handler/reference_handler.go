package handler

import (
	"encoding/json"
	"net/http"

	"github.com/practice-cloud-run-functions/myapp/reference-api/service"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.FetchUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}