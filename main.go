package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Define the directory where the logos are stored
	logoDir := "./logos"

	// Check if directory exists
	if _, err := os.Stat(logoDir); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist. Please create it and add some images.\n", logoDir)
		return
	}

	// Handle the image serving
	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		// Extract the image file name from the URL
		imageName := r.URL.Path[len("/images/"):]

		// Build the full file path
		filePath := filepath.Join(logoDir, imageName)

		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "Image not found", http.StatusNotFound)
			return
		}

		// Serve the file
		http.ServeFile(w, r, filePath)
	})

	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
