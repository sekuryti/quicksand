package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/fogleman/gg"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run script.go /path/to/your/directory num_threads --videos|--images")
		return
	}

	directory := os.Args[1]
	numThreads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid num_threads argument:", err)
		return
	}

	fileType := strings.ToLower(os.Args[3])
	if fileType != "--videos" && fileType != "--images" {
		fmt.Println("Invalid file type argument. Use --videos or --images.")
		return
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, numThreads)

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(directory, file.Name())
			semaphore <- struct{}{} // Acquire a semaphore
			wg.Add(1)
			go func(filePath string) {
				defer wg.Done()
				defer func() {
					<-semaphore // Release the semaphore
				}()

				if fileType == "--videos" {
					// Check if it's a video
					if isVideo(filePath) {
						cmd := exec.Command("ffmpeg", "-v", "error", "-i", filePath, "-f", "null", "-")
						err := cmd.Run()
						if err != nil {
							fmt.Printf("Deleting corrupted video: %s\n", filePath)
							os.Remove(filePath)
						} else {
							fmt.Printf("Video is OK: %s\n", filePath)
						}
					}
				} else if fileType == "--images" {
					// Check if it's an image
					if isImage(filePath) {
						if !isImageValid(filePath) {
							fmt.Printf("Deleting corrupted image: %s\n", filePath)
							os.Remove(filePath)
						} else {
							fmt.Printf("Image is OK: %s\n", filePath)
						}
					}
				}
			}(filePath)
		}
	}

	wg.Wait()
}

// Function to check if a file is a video (based on file extension)
func isVideo(filePath string) bool {
	videoExtensions := []string{".mp4", ".avi", ".mkv", ".mov", ".flv", ".wmv", ".webm"}
	ext := filepath.Ext(filePath)
	for _, videoExt := range videoExtensions {
		if strings.EqualFold(ext, videoExt) {
			return true
		}
	}
	return false
}

// Function to check if a file is an image (based on file extension)
func isImage(filePath string) bool {
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	ext := filepath.Ext(filePath)
	for _, imageExt := range imageExtensions {
		if strings.EqualFold(ext, imageExt) {
			return true
		}
	}
	return false
}

// Function to check if an image is valid
func isImageValid(filePath string) bool {
	img, err := gg.LoadImage(filePath)
	return err == nil && img != nil
}
