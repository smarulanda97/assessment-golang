package web

import (
	"context"
	"log"
	"os"

	"github.com/smarulanda97/assessment-golang/internal/server"
)

// StartApiServer run server with the API
func StartServerWeb() {
	APP_PORT := os.Getenv("APP_PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")

	config := &server.Config{
		Port:      APP_PORT,
		JWTSecret: JWT_SECRET,
	}

	serverApi, err := server.NewServer(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}

	if err := serverApi.Start(BindRoutesWeb); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("server have started on port %s", APP_PORT)
		log.Println("...")
	}
}
