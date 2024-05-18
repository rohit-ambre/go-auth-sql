package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rohit-ambre/go-auth-sql/controllers"
)

func TokenValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside middleware")

		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			resp := controllers.Response{Success: false, Msg: "Auth token is missing"}
			json.NewEncoder(w).Encode(resp)
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			resp := controllers.Response{Success: false, Msg: "Something went wrong"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		fmt.Println("token", token)

		if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			if float64(time.Now().Unix()) > claim["exp"].(float64) {
				resp := controllers.Response{Success: false, Msg: "Token is expired"}
				json.NewEncoder(w).Encode(resp)
				return
			}
			// how to set this into req context
			fmt.Println("UserID", claim["userId"])
			next.ServeHTTP(w, r)
		} else {
			resp := controllers.Response{Success: false, Msg: "Token is not valid"}
			json.NewEncoder(w).Encode(resp)
			return
		}
	})
}
