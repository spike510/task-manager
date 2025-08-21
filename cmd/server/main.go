package main

import (
	"log"

	"github.com/spike510/task-manager/internal/http"
)

func main() {
	r := http.NewRouter()

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
