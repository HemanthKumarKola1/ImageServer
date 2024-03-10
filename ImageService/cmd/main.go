package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/hemanthkumarkola1/imageservice/api"
	db "github.com/hemanthkumarkola1/imageservice/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PageData represents data to be passed to HTML template
type PageData struct {
	Images []string
}

func main() {
	// Define HTTP routes

	ctx := context.Background()
	connPool, err := pgxpool.New(ctx, "postgresql://root1:secret1@localhost:5432/schedules?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	q := db.New(connPool)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) { api.UploadImageHandler(w, r, q) })
	http.HandleFunc("/get-image", func(w http.ResponseWriter, r *http.Request) { api.GetImageHandler(w, r, q) })

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Render the HTML template
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
