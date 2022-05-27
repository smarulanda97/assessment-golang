package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/smarulanda97/assessment-golang/cmd/api"
	"github.com/smarulanda97/assessment-golang/cmd/web"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("an error have ocurred loading .env file on api server")
	}

	// Initialize a server of API (Backend) and APP (Frontend)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		web.StartServerWeb()
		wg.Done()
	}()
	go func() {
		api.StartServerApi()
		wg.Done()
	}()

	wg.Wait()
}
