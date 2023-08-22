package main

import (
	"article-tags/internal/database/connection"
	"article-tags/internal/handler"
	"article-tags/internal/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load env config
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("not able to load env", err)
	}

	conn, err := connection.GetConnection()
	if err != nil {
		log.Fatalln("db connect error", err)
	}

	app := handler.NewApplication(conn)

	// register routes
	r := routes.RegisterRoutes(app)

	// run server
	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("run error:", err)
	}
}
