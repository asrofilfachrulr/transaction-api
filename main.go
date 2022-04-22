package main

import (
	"flag"
	"log"
	"os"

	"github.com/asrofilfachrulr/transaction-api/config"
	"github.com/asrofilfachrulr/transaction-api/docs"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		// exit program in fail load .env
		log.Fatalln(err.Error())
	}

	// set schemes only http if run in local
	docs.SwaggerInfo.Schemes = []string{"http"}

	// set debug mode by given flag
	debugMode := flag.String("d", "", "")
	// drop all table before migration
	drop := flag.String("drop", "", "")
	flag.Parse()

	// set to this runtime env
	os.Setenv("DEBUG_MODE", *debugMode)
	os.Setenv("DROP", *drop)

	// get host url from env
	docs.SwaggerInfo.Host = "localhost:8080"
}

// @title Transaction API
// @version 1.0
// @description API provide simple transaction

// @contact.name Developer
// @contact.email riidloa@gmail.com

// @BasePath /api/v1
func main() {
	server := config.NewServer()
	server.Init()
	server.Run()
}
