package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"eventhub/middleware"
	"eventhub/services"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserContextKey).(int)

	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)

	err := services.CreateEvent(
		data["name"],
		data["location"],
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("Event Created"))
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	events, _ := services.GetEvents()
	json.NewEncoder(w).Encode(events)
}

func JoinEvent(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserContextKey).(int)

	eventIDStr := r.URL.Query().Get("event_id")
	eventID, _ := strconv.Atoi(eventIDStr)

	err := services.JoinEvent(userID, eventID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("Joined Event"))
}
