package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // Import png to allow decoding PNG files
	"net/http"
	"os"
)

func main() {
	// 1. Setup the route (Endpoint)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// If the request is GET, show the upload form (for testing purposes)
		if r.Method != "POST" {
			showUploadForm(w)
			return
		}

		// 2. Parse the uploaded file
		// "image" is the name of the input field in the form
		file, _, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Failed to read the file.", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// 3. Decode: Convert the file into image object (pixels)
		img, _, err := image.Decode(file)
		if err != nil {
			http.Error(w, "Invalid image format. Please upload JPEG or PNG.", http.StatusBadRequest)
			return
		}

		// 4. Encode & Compress: Re-draw the image with lower quality
		// Quality ranges from 1 to 100.
		// Lower quality = Smaller file size. 30-50 is usually a good balance.
		w.Header().Set("Content-Type", "image/jpeg")
		err = jpeg.Encode(w, img, &jpeg.Options{Quality: 40})
		if err != nil {
			http.Error(w, "Failed to compress image.", http.StatusInternalServerError)
			return
		}
	})

	// Get the PORT from the environment variable (Required for Cloud Hosting)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port for localhost
	}

	fmt.Println("Image Compressor Service is running on port:", port)
	http.ListenAndServe(":"+port, nil)
}

// Function to display a simple HTML form for testing
func showUploadForm(w http.ResponseWriter) {
	html := `
	<html>
		<head><title>Image Compressor</title></head>
		<body>
			<h2>Upload Image to Compress</h2>
			<form action="/" method="post" enctype="multipart/form-data">
				<input type="file" name="image" accept="image/*">
				<br><br>
				<input type="submit" value="Compress Now">
			</form>
		</body>
	</html>
	`
	w.Write([]byte(html))
}
