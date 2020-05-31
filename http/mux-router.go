package router

import (
	"fmt"
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
	handler       = c.Handler(muxDispatcher)
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET", "OPTIONS")
	handler = c.Handler(muxDispatcher)
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST", "OPTIONS")
	handler = c.Handler(muxDispatcher)
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(":"+port, handler)
}
