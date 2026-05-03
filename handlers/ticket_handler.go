package handlers

import (
	"encoding/json"
	"eventhub/middleware"
	"eventhub/services"
	"net/http"
	"strconv"
	"time"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 🔐 Get user from context (if needed)
	userID := r.Context().Value(middleware.UserContextKey).(int)
	_ = userID // use later for organizer validation

	// 🔁 Convert values safely
	eventID, err := strconv.Atoi(data["event_id"])
	if err != nil {
		http.Error(w, "Invalid event_id", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(data["price"], 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	totalQty, err := strconv.Atoi(data["total_quantity"])
	if err != nil {
		http.Error(w, "Invalid total_quantity", http.StatusBadRequest)
		return
	}

	// ❗ remainingQty should NOT come from user
	remainingQty := totalQty

	saleStart, err := time.Parse(time.RFC3339, data["sale_start_time"])
	if err != nil {
		http.Error(w, "Invalid sale_start_time", http.StatusBadRequest)
		return
	}

	saleEnd, err := time.Parse(time.RFC3339, data["sale_end_time"])
	if err != nil {
		http.Error(w, "Invalid sale_end_time", http.StatusBadRequest)
		return
	}

	// 🚀 Call service
	err = services.CreateTicket(
		eventID,
		data["name"],
		data["description"],
		totalQty,
		remainingQty,
		price,
		saleStart,
		saleEnd,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ticket Created"))
}
