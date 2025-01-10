package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	fmt.Println("Enter the local path of the PDF file:")
	var filePath string
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Open the PDF file.
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Get PDF information using pdfcpu.
	info, err := api.PDFInfo(file, filePath, nil, nil) // nil for default configuration and all pages.
	if err != nil {
		log.Fatalf("Failed to extract PDF info: %v", err)
	}

	// Display PDF information.
	fmt.Println("PDF Information:")
	fmt.Printf("File Name: %s\n", filePath)
	fmt.Printf("Number of Pages: %d\n", info.PageCount)
	fmt.Printf("Title: %s\n", info.Title)
	fmt.Printf("Author: %s\n", info.Author)
	fmt.Printf("Subject: %s\n", info.Subject)
	fmt.Printf("Producer: %s\n", info.Producer)
	fmt.Printf("Creator: %s\n", info.Creator)
	fmt.Printf("Creation Date: %s\n", info.CreationDate)
	fmt.Printf("Modification Date: %s\n", info.ModificationDate)

	// Display page dimensions.
	fmt.Println("\nUnique Page Dimensions:")
	for dim := range info.PageDimensions {
		fmt.Printf("Width: %.2f, Height: %.2f\n", dim.Width, dim.Height)
	}

	outputFile := "output.md"
	cmd := exec.Command("pdftotext", filePath, outputFile)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to extract text: %v", err)
	}

	fmt.Printf("Text extracted successfully! Check the file: %s\n", outputFile)
}
