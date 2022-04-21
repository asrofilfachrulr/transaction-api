package main

import (
	"flag"
	"log"
	"os"

	"github.com/asrofilfachrulr/transaction-api/config"
	"github.com/joho/godotenv"
	"github.com/swaggo/swag/example/basic/docs"
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

func main() {
	server := config.NewServer()
	server.Init()
	server.Run()
}
