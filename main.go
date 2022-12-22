package main

import (
	"HashimJVZ/image-upload/auto"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	maxFileSize = 5 << 20 // <<20 indicates MB to bytes

)

var (
	server_address = "http://localhost"
	port           = ":8081"
)

// return true if the extension is in allowed extension list.
func isAllowedExtension(fileExtension string) bool {
	allowedExtensions := []string{".png", ".jpg", ".jpeg", ".bmp", ".tiff", ".tif"}
	for _, e := range allowedExtensions {
		if e == fileExtension {
			return true
		}
	}
	return false
}

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "uploading...")
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		fmt.Fprintln(w, "Upload Failed")
		return
	}

	// retrieving file from html form
	file, fileheader, err := r.FormFile("imgFile")
	if err != nil {
		fmt.Fprintln(w, "Error Retrieving the File")
		fmt.Fprintln(w, err)
		return
	}
	defer file.Close()

	// Check file size
	if fileheader.Size > maxFileSize {
		fmt.Fprintln(w, "File size is greater than 5MB")
		fmt.Fprintln(w, "Upload Failed")
		return
	}

	// check file extension
	fileExtension := filepath.Ext(fileheader.Filename)
	if !isAllowedExtension(fileExtension) {
		fmt.Fprintln(w, "File extension not supported")
		fmt.Fprintln(w, "Upload Failed")
		return
	}

	// creating temp file
	pattern := fmt.Sprintf("*%s", fileExtension)
	directory := "static/images"
	tempFile, err := os.CreateTemp(directory, pattern)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	//show upload success
	fmt.Fprintln(w, "Upload Succesful")

	filename := tempFile.Name()[14:]

	log.Printf("%s uploaded successfully", filename)

	// show preview link
	fmt.Fprintf(w, "Image link: %s%s/images/%s\n", server_address, port, filename)
}

func main() {

	//loading environment file(.env by default)
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// loading port from environment file
	server_address = os.Getenv("SERVER_ADDRESS")
	if server_address == "" {
		log.Fatal("Error loading SERVER_ADDRESS from .env file")
	}

	port = os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error loading port from .env file")
	}

	go auto.DeleteOld(7)

	http.HandleFunc("/uploadimage", uploadImageHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.ListenAndServe(port, nil)
}
