package tago

import (
	"log"
	"os"
)

func Build(page *Element, path string) {
	f, err := os.Create(path + ".html")
	if err != nil {
		log.Fatalf("failed to create file : %v", err)
	}
	defer f.Close()

	if err := page.Render(f, 0); err != nil {
		log.Fatalf("failed to render HTML : %v", err)
	}
}
