package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	// Load the image of the camel
	img := gocv.IMRead("camel.jpg", gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Error: Could not load image")
		os.Exit(1)
	}

	// Load the classifier for detecting camels
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("camel_classifier.xml") {
		fmt.Println("Error: Could not load classifier")
		os.Exit(1)
	}

	// Detect camels in the image
	rects := classifier.DetectMultiScale(img)

	// Check if a camel was detected with at least 70% certainty
	if len(rects) > 0 && rects[0].Confidence > 0.7 {
		fmt.Println("Camel detected")
	} else {
		fmt.Println("Camel not detected")
	}
}
