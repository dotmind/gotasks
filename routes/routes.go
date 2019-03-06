package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"../config"
)

type Route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type EmptyResponse struct{}

type Response struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

var appConfig = config.Load()

var routes = Routes{
	Route{
		Path:        "/",
		Method:      "GET",
		HandlerFunc: RootHandler,
	},
	Route{
		Path:        "/add",
		Method:      "GET", // @TODO Only for dev
		HandlerFunc: AddHandler,
	},
	Route{
		Path:        "/getall",
		Method:      "GET",
		HandlerFunc: GetAllHandler,
	},
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := isOwnNetwork()

		if auth == false {
			jsonFailResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Handler(route.HandlerFunc)
	}

	router.Use(authMiddleware)

	return router
}
