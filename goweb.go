package main

import (
	"encoding/gob"
	"log"

	"github.com/DotsCore/go-web/config"
	"github.com/DotsCore/go-web/database/model"
	"github.com/DotsCore/go-web/register"
	"github.com/joho/godotenv"
)

// Main Go-Web entry point.
func main() {
	gob.Register(&model.User{})

	entities := register.BaseEntities()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	foundation.Start(entities, config.GetSever())
}
