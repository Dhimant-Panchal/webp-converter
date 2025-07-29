package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/webp"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <inputFile.webp> <outputFile.png|jpg|jpeg>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Opening Webp
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Could not open input file %s:\n%v", inputFile, err)
	}
	defer file.Close()

	// Decoding Webp
	img, err := webp.Decode(file)
	if err != nil {
		log.Fatalf("Could not decode WebP image: %v", err)
	}

	// Creating output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Could not create output file %s:\n%v", outputFile, err)
	}
	defer outFile.Close()

	// Encode the image to the desired format
	ext := strings.ToLower(filepath.Ext(outputFile))
	switch ext {
	case ".png":
		err = png.Encode(outFile, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
	default:
		fmt.Printf("Unsupported output format: %s. Please use .png or .jpg\n", ext)
		return
	}

	if err != nil {
		log.Fatalf("Could not encode to output file: %v", err)
	}

	fmt.Printf("Converted %s to %s\n", inputFile, outputFile)

}
