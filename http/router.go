package router

import "net/http"

type Router interface {
	GET(uri string, port string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, port string, f func(w http.ResponseWriter, r *http.Request))
}
