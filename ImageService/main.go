package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

// PageData represents data to be passed to HTML template
type PageData struct {
	Images []string
}

func main() {
	// Define HTTP routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadImageHandler)
	http.HandleFunc("/get-images", getImageHandler)

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

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read user id and schedule time
	userID := r.FormValue("userID")
	scheduleTimeString := r.FormValue("futureTime")

	if userID == "" || scheduleTimeString == "" {
		http.Error(w, "Error in reading request header", http.StatusInternalServerError)
		return
	}

	nameOfTheFile := "images/" + userID + "$" + scheduleTimeString + ".jpg" // Destination location

	// Read the request body
	imageBlob, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// You can now process the image blob as needed
	// For example, you can save it to a file, store it in a database, or perform other operations

	// Example: Saving the image blob to a file
	err = os.WriteFile(nameOfTheFile, imageBlob, 0644)
	if err != nil {
		http.Error(w, "Error saving image", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Image uploaded successfully")
}

var imageDirectory = "./images/" // Directory containing the images

func getImageHandler(w http.ResponseWriter, r *http.Request) {
	// Get the current time
	// currentTime := time.Now()

	// Iterate through the images in the directory
	// files, err := os.ReadDir(imageDirectory)
	// if err != nil {
	// 	http.Error(w, "Error reading image directory", http.StatusInternalServerError)
	// 	return
	// }

	// Find the image whose filename corresponds to the current time
	var imageName = "kola$2024-03-03T11:16.jpg"
	// Serve the image
	http.ServeFile(w, r, imageDirectory+imageName)
	// for _, file := range files {
	// 	// Skip directories
	// 	if file.IsDir() {
	// 		continue
	// 	}

	// 	// Extract the time from the filename
	// 	fileNameParts := strings.Split(file.Name(), ".")

	// 	imageTime := strings.Split(fileNameParts[0], "$")[1]

	// 	// Parse the time from the filename
	// 	parsedTime, err := time.Parse("2024-03-03T11:16", imageTime) // Assuming filename format: "2006-01-02_15-04-05.jpg"
	// 	if err != nil {
	// 		continue
	// 	}

	// 	// If the image time matches the current time, set imageName and break the loop
	// 	if parsedTime.Equal(currentTime) {
	// 		imageName = file.Name()
	// 		break
	// 	}
	// }

	// // If imageName is empty, no image found for the current time
	// if imageName == "" {
	// 	http.Error(w, "No image found for the current time", http.StatusNotFound)
	// 	return
	// }

}
