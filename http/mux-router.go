package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type muxRouter struct{}

var (
	c = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
)
var (
	muxDispatcher = mux.NewRouter()
	methodType    = ""
	handler       = c.Handler(muxDispatcher)
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET", "OPTIONS")
	methodType = "Get"
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST", "OPTIONS")
	methodType = "Post"
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	if methodType == "Get" {
		http.ListenAndServe(":"+port, muxDispatcher)
	} else {
		log.Println("post method")
		http.ListenAndServe(":"+port, muxDispatcher)
	}
}
