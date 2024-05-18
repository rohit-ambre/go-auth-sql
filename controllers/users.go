package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rohit-ambre/go-auth-sql/models"
)

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := models.GetAllUsers()
	if err != nil {
		resp := Response{Success: false, Msg: "Something went wrong"}
		json.NewEncoder(w).Encode(resp)
	}

	resp := Response{Success: true, Msg: "Success", Data: res}
	json.NewEncoder(w).Encode(resp)
}
