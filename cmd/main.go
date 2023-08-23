package main

import (
	"article-tags/internal/database/connection"
	"article-tags/internal/handler"
	"article-tags/internal/routes"
	"context"
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

	// init db setup create tables
	err = initSetup(app)
	if err != nil {
		log.Fatalln("db setup error", err)
	}

	// register routes
	r := routes.RegisterRoutes(app)

	// run server
	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("run error:", err)
	}
}

func initSetup(app *handler.Application) error {
	ctx := context.Background()
	// check table if exists or not. if not present create new
	err := app.ArticleStore.DescribeTable(ctx)
	if err != nil {
		err = app.ArticleStore.CreateTable(ctx)
		if err != nil {
			log.Println("create table failed")
			return err
		}
	}

	return nil
}
