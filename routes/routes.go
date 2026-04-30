package routes

import (
	"net/http"

	"eventhub/handlers"
	"eventhub/middleware"
)

func RegisterRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("EventHub API Running 🚀"))
	})

	// Auth
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	// Protected
	http.HandleFunc("/events/create", middleware.AuthMiddleware(handlers.CreateEvent))
	http.HandleFunc("/events", handlers.GetEvents)
	http.HandleFunc("/events/join", middleware.AuthMiddleware(handlers.JoinEvent))
}
