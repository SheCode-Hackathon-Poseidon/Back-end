package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sample/api/responses"
)

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data, including the uploaded file
	err := r.ParseMultipartForm(10 << 20) // Set the maximum file size in bytes (10 MB in this case)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("file") // "file" is the field name for the uploaded file
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Specify the folder path where you want to save the image
	saveFolder := "./image" // Change this path to your desired folder

	// Ensure the folder exists; create it if it doesn't
	if err := os.MkdirAll(saveFolder, os.ModePerm); err != nil {
		http.Error(w, "Unable to create folder for saving the image", http.StatusInternalServerError)
		return
	}

	// Construct the file path to save the uploaded image
	savePath := filepath.Join(saveFolder, handler.Filename)

	// Create a new file to save the uploaded image
	outFile, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "Unable to create file for saving the image", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Copy the file data to the new file
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Unable to save the image", http.StatusInternalServerError)
		return
	}

	imageUrl := handler.Filename

	responses.JSON(w, http.StatusOK, map[string]interface{}{
		"exitcode": 0,
		"message":  "Image uploaded and saved as",
		"url":      imageUrl,
	})
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	// Parse the image filename from the URL
    filename := r.URL.Path[1:] // Remove the leading forward slash
	fmt.Printf("Get image file name = %s\n", filename)
    
    // Open the image file
    file, err := os.Open(filename)
    if err != nil {
        http.Error(w, "Image not found", http.StatusNotFound)
        return
    }
    defer file.Close()

    // Set the content type header (change this to match your image format)
    w.Header().Set("Content-Type", "image/*")

    // Copy the image data to the response writer
    _, err = io.Copy(w, file)
    if err != nil {
        http.Error(w, "Failed to copy image data", http.StatusInternalServerError)
        return
    }
}
