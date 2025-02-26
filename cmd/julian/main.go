package main

import (
	"log"
	"os"

	foundation "github.com/DotsCore/go-web-framework"
	"github.com/DotsCore/go-web/register"
	"github.com/joho/godotenv"
)

func main() {
	entities := register.BaseEntities()
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found, some features may not work!")
	}

	foundation.StartCommand(os.Args[1:], entities)
}
