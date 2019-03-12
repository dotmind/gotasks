package routes

import (
	"encoding/json"
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

func (r Response) WithSuccess() Response {
	r.Success = true
	return r
}

func (r Response) WithError() Response {
	r.Success = false
	return r
}

func (r Response) WithPayload(payload interface{}) Response {
	r.Payload = payload
	return r
}

func (r Response) Send(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(r)
	return
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
		Method:      "POST",
		HandlerFunc: AddHandler,
	},
	Route{
		Path:        "/update",
		Method:      "POST",
		HandlerFunc: UpdateHandler,
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
			Response{}.WithError().Send(w)
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
