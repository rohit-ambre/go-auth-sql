package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/rohit-ambre/go-auth-sql/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println(os.Getenv("PORT"))

	server := http.Server{
		Addr:    "localhost:" + os.Getenv("PORT"),
		Handler: routes.Init(),
	}
	fmt.Println("Server listening on port :" + os.Getenv("PORT"))
	server.ListenAndServe()
}
