package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")
	if v == "" {
		fmt.Fprintf(w, "No video id provided")
		return
	}

	err := downloadVideoExtractAudio(v, w)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

}

func downloadVideoExtractAudio(id string, out io.Writer) error {
	url := fmt.Sprintf("https://www.youtube.com/watch?v=" + id)

	r, w := io.Pipe()

	defer r.Close()

	ytdl := exec.Command("youtube-dl", url, "-o-")

	ytdl.Stdout = w
	ytdl.Stderr = os.Stderr

	ffmpeg := exec.Command("ffmpeg", "-i", "/dev/stdin", "-f", "mp3", "-ab", "96000", "-vn", "-")

	ffmpeg.Stdin = r
	ffmpeg.Stdout = out

	go ytdl.Run()

	err := ffmpeg.Run()

	return err

}

func main() {

	port := os.Getenv("PORT")

	if port == "" {

		port = "8080"

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Welcome to Youtube Audio Extractor\n\nUsage: /watch?v=VIDEO_ID\n\n")

	})

	http.HandleFunc("/watch", handler)

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
