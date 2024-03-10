package api

import (
	"net/http"

	db "github.com/hemanthkumarkola1/imageservice/db/sqlc"
)

func GetImageHandler(w http.ResponseWriter, r *http.Request, q *db.Queries) {

	schedule, err := q.GetSchedule(r.Context(), 12)
	if err != nil {
		http.Error(w, "Unable to get image data from database", http.StatusInternalServerError)
		return
	}

	// Set the content type header to indicate that the response contains an image
	w.Header().Set("Content-Type", "image/jpeg") // Change the content type based on the image type

	// Write the image data to the response
	w.Write(schedule.ImageData)
}
