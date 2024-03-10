package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	db "github.com/hemanthkumarkola1/imageservice/db/sqlc"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request, q *db.Queries) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read user id and schedule time
	userID := r.FormValue("userID")
	scheduleTimeString := r.FormValue("futureTime")
	message := r.FormValue("message")

	if userID == "" || scheduleTimeString == "" {
		http.Error(w, "Error in reading request header", http.StatusInternalServerError)
		return
	}

	// Read the request body
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}

	imageData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	intTime, _ := strconv.Atoi(scheduleTimeString)

	_, err = q.CreateSchedule(r.Context(), db.CreateScheduleParams{int32(intTime), userID, message, imageData})
	if err != nil {
		http.Error(w, "Unable to write to db", http.StatusInternalServerError)
	}
	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Image uploaded successfully")
}
