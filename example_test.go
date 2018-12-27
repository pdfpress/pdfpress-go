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
