package routes

import (
	"fmt"
	"net/http"

	"github.com/rohit-ambre/go-auth-sql/controllers"
	"github.com/rohit-ambre/go-auth-sql/middleware"
)

// func handle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("reached here")
// 	w.Write([]byte("response"))
// }

// func handleSignup(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("handleSignup")
// 	w.Write([]byte("handleSignup"))
// }

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}

func Init() *http.ServeMux {
	fmt.Println("inside route Init")
	router := http.NewServeMux()

	// router.HandleFunc("/some", handle)

	// router.HandleFunc("GET /all", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("all route")
	// 	w.Write([]byte("all"))
	// })

	// router.HandleFunc("GET /obj", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("obj route")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	res := Response{Success: true, Msg: "some", Data: ""}
	// 	json.NewEncoder(w).Encode(res)
	// 	// w.Write(Response{success: true, msg: "some", data: {hi: "msg"}})
	// })

	router.HandleFunc("POST /auth/signup", controllers.HandleSignup)
	router.HandleFunc("POST /auth/login", controllers.HandleLogin)

	apiRouter := http.NewServeMux()
	apiRouter.Handle("/api/", http.StripPrefix("/api", router))

	authRouter := http.NewServeMux()
	// need api as prefix
	authRouter.HandleFunc("GET /users/getUsers", controllers.HandleGetUsers)

	apiRouter.Handle("/", middleware.TokenValidator(authRouter))

	return apiRouter
}
