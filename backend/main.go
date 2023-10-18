package main

import (
	"log"
	"todo/api/routes"
	"todo/infrastructure"

	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	infrastructure.ConnectToDB()
}

func main(){
	r := routes.RouterSetup()
	r.Run()
}