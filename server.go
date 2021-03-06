package main

import (
    "fmt"
	"net/http"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"errors"
	"log"
	"os"
	"twitter-hangouts/routes"
	"github.com/gorilla/handlers"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func main() {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	port, port_exists := os.LookupEnv("PORT")

	if !port_exists {
		port = "8000"
	}

	var handler http.Handler
	{
    	handler = handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "Authorization", "Content-Type"}),
			handlers.ExposedHeaders([]string{""}),
			handlers.MaxAge(10),
			handlers.AllowCredentials(),
    	)(routes.Init())
    	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
	}
	http.Handle("/", handler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log(err)
		log.Fatal(err)
	}

}