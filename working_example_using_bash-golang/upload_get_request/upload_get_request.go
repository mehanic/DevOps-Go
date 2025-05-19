package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
	"io"
)

const (
	server = "10.230.229.11"
	port   = 8080
)

// getFilename generates a random filename of 12 alphanumeric characters.
func getFilename() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	fn := make([]byte, 12)
	for i := range fn {
		fn[i] = chars[rand.Intn(len(chars))]
	}
	return string(fn)
}

// uploadHandler handles file upload requests.
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form to get file data
		err := r.ParseMultipartForm(10 << 20) // Limit upload size to 10MB
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the file from form
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to get file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create a new file with a random name
		fn := getFilename()
		out, err := os.Create(fn)
		if err != nil {
			http.Error(w, "Unable to create file", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		// Write the uploaded file data to the new file
		if _, err := io.Copy(out, file); err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "File uploaded successfully!")
		return
	}

	// For GET requests, display the upload form
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `
		<h1>Upload a File</h1>
		<form action="/" method="post" enctype="multipart/form-data">
			<input type="file" name="file" placeholder="Enter a filename."></input><br />
			<input type="submit" value="Import">
		</form>
	`)
}

func main() {
	http.HandleFunc("/", uploadHandler)
	fmt.Printf("Starting server on %s:%d, use <Ctrl-C> to stop\n", server, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", server, port), nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}

