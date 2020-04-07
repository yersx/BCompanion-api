// main
package main

import (
	"bcompanion/controller"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/users/authorize", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/users/auth", controller.AuthHandler).Methods("GET")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	r.HandleFunc("/users/getUser", controller.ProfileHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
