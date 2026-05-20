package main

import (
	"log"
	"tech-tracker-go/api/routes"
)

func main() {
	appRouter := routes.SetupApp()

	log.Println("Server is running on port 8080 🚀")


	err := appRouter.Run(":3000")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
