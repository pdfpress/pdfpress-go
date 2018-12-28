package pdfpress_test

import (
	"context"
	"github.com/pdfpress/pdfpress-go"
	"log"
)

func Example_generateFromHTML() {
	ctx := context.Background()
	pdfpress.Key = "api_key"
	params := &pdfpress.GenerateParams{
		HTML: "<html />",
	}
	pdf, err := pdfpress.Generate(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	defer pdf.Close()
	log.Println("Generated pdf")
}

func Example_generateFromURL() {
	ctx := context.Background()
	pdfpress.Key = "api_key"
	params := &pdfpress.GenerateParams{
		URL: "https://pdfpress.app",
	}
	pdf, err := pdfpress.Generate(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	defer pdf.Close()
	log.Println("Generated pdf")
}

func Example_generateFromTemplate() {
	ctx := context.Background()
	templateParams := map[string]string{
		"username": "johnsmith",
	}
	templateID := "uuid-uuid-uuid-uuid"
	pdf, err := pdfpress.GenerateFromTemplate(ctx, templateID, templateParams)
	if err != nil {
		log.Fatal(err)
	}
	defer pdf.Close()
	log.Println("Generated pdf")
}
