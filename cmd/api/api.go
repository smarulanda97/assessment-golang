package api

import (
	"context"
	"log"
	"os"

	"github.com/smarulanda97/assessment-golang/internal/server"
)

// StartApiServer run server with the API
func StartServerApi() {
	API_PORT := os.Getenv("API_PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")

	config := &server.Config{
		Port:      API_PORT,
		JWTSecret: JWT_SECRET,
	}

	serverApi, err := server.NewServer(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}

	if err := serverApi.Start(BindRoutesApi); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("server have started on port %s", API_PORT)
		log.Println("...")
	}
}
