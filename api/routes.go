package api

import (
	"github.com/gorilla/mux"

	"github.com/alifhaider/carehub-backend/api/handlers"
	"github.com/alifhaider/carehub-backend/api/middleware"
)

func InitRoutes() *mux.Router {
    r := mux.NewRouter()

    // Use middleware for logging requests
    r.Use(middleware.LoggingMiddleware)

    // Define routes and handlers
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")

		r.HandleFunc("/doctors", handlers.GetDoctors).Methods("GET")
    r.HandleFunc("/doctors", handlers.CreateDoctor).Methods("POST")
    r.HandleFunc("/doctors", handlers.UpdateDoctor).Methods("PUT")
    r.HandleFunc("/doctors", handlers.DeleteDoctor).Methods("DELETE")

    return r
}