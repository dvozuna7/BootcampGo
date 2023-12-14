package pkg

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"practiva/web/internal/domain"
)

func FullfilDB(path string) *[]domain.Product {
	data, err := os.Open(path)
	fullPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Full path: ", fullPath, " --- Error", err)
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	var slice []domain.Product
	err = json.Unmarshal(dataRead, &slice)
	if err != nil {
		log.Fatal(err)
	}
	return &slice
}
