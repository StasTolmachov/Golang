package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <pdf-file>")
		os.Exit(1)
	}

	pdfPath := os.Args[1]
	err := extractTextFromPDF(pdfPath)
	if err != nil {
		fmt.Printf("Ошибка при извлечении текста из файла %s: %v\n", pdfPath, err)
	}
}

func extractTextFromPDF(pdfPath string) error {
	// Открыть PDF-файл.
	f, err := os.Open(pdfPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Загрузить PDF-документ.
	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return err
	}

	// Обход страниц и извлечение текста.
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return err
		}

		// Создание текстового извлекателя.
		textExtractor, err := extractor.New(page)
		if err != nil {
			return err
		}

		pageText, err := textExtractor.ExtractText()
		if err != nil {
			return err
		}

		fmt.Printf("Текст страницы %d:\n%s\n", i, pageText)
	}

	return nil
}
