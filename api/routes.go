package api

import (
	"github.com/gorilla/mux"

	"github.com/alifhaider/carehub-backend/api/handlers"
)

func InitRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    return r
}
