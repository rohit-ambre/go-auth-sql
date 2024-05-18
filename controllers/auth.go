package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rohit-ambre/go-auth-sql/models"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleSignup")
	w.Header().Set("Content-Type", "application/json")

	var reqBody models.SignUpReq

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if reqBody.EmailID == "" {
		resp := Response{Success: false, Msg: "EmailID is required"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	user, err := models.GetUserByEmail(reqBody.EmailID)
	if err != nil {
		fmt.Println("error in auth", err)
		resp := Response{Success: false, Msg: "Something went wrong"}
		json.NewEncoder(w).Encode(resp)
	} else if user.UserID != 0 {
		resp := Response{Success: false, Msg: "User Already exists"}
		json.NewEncoder(w).Encode(resp)
	} else {

		hashedPw, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)

		if err != nil {
			resp := Response{Success: false, Msg: "Something went wrong"}
			json.NewEncoder(w).Encode(resp)
		}
		fmt.Println("hashed", string(hashedPw))
		userReq := models.User{EmailID: reqBody.EmailID, Password: string(hashedPw), FirstName: reqBody.FirstName, LastName: reqBody.LastName, Active: true}

		userRes, err := models.CreateUser(userReq)
		if err != nil {
			fmt.Println("Error creating user")
			resp := Response{Success: false, Msg: "Something went wrong"}
			json.NewEncoder(w).Encode(resp)
		}
		fmt.Println("userRes", userRes)
		json.NewEncoder(w).Encode(Response{Success: true, Msg: "user created", Data: userRes})
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HandleLogin")
	w.Header().Set("Content-Type", "application/json")

	var reqBody models.LoginReq

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if reqBody.EmailID == "" {
		resp := Response{Success: false, Msg: "EmailID is required"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	user, err := models.GetUserByEmail(reqBody.EmailID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.UserID == 0 {
		resp := Response{Success: false, Msg: "User Not found"}
		json.NewEncoder(w).Encode(resp)
	}

	err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))

	if err1 != nil {
		resp := Response{Success: false, Msg: "Invalid Email or password"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	// resp := Response{Success: true, Msg: "Validated"}
	// json.NewEncoder(w).Encode(resp)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.UserID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println("Error creating token", err)
		resp := Response{Success: false, Msg: "Error creating token"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(Response{Success: true, Msg: "Token generated", Data: tokenString})
}
