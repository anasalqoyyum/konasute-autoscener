package utils

import (
	"image"
	"log"
	"path/filepath"
	"sdvx-autoscener/pkg/constant"

	"gocv.io/x/gocv"
)

// func containsObject(img gocv.Mat) bool {
func ContainsObject() bool {
	img := gocv.IMRead(filepath.Join(constant.OUTPUT_PATH, "screen.png"), gocv.IMReadColor)
	if img.Empty() {
		log.Println("Error loading output image")
		return false
	}
	defer img.Close()

	// Resize the image to 1080x1920 pixels
	// for consistency's sake, considering konasute should mostly run at 1080p vertical
	newSize := image.Point{X: 1080, Y: 1920}
	resizedImg := gocv.NewMat()
	defer resizedImg.Close()

	gocv.Resize(img, &resizedImg, newSize, 0, 0, gocv.InterpolationLinear)

	template := gocv.IMRead(filepath.Join(constant.RESOURCES_PATH, "song_info.png"), gocv.IMReadGrayScale)
	if template.Empty() {
		log.Println("Error loading template image")
		return false
	}
	defer template.Close()

	// Convert the captured frame to grayscale
	grayImg := gocv.NewMat()
	defer grayImg.Close()

	gocv.CvtColor(resizedImg, &grayImg, gocv.ColorBGRToGray)

	// Perform template matching
	result := gocv.NewMat()
	defer result.Close()

	mask := gocv.NewMat()
	defer mask.Close()

	gocv.MatchTemplate(grayImg, template, &result, gocv.TmCcoeffNormed, mask)

	// Check if the matching result exceeds a threshold
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	threshold := float32(0.8)
	if maxVal >= threshold {
		log.Printf("Object found at location: %v with confidence: %f\n", maxLoc, maxVal)
		return true
	}

	return false
}
