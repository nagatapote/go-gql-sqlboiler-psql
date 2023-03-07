package main

import (
	"go-gql-sqlboiler-psql/config"
	"go-gql-sqlboiler-psql/infrastructure/db"
	"go-gql-sqlboiler-psql/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const defaultPort = "8000"

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		envLoad()
	}

	if err := envconfig.Process("", &config.Conf); err != nil {
		panic(err)
	}
	db.Init()

	e := gin.Default()

	router.HandlerGraphql(e)

	e.GET("/",
		router.HandlerPlayGround(),
	)

	log.Printf("connect to :%s/ for GraphQL playground", port)
	e.Run(":" + port)
}
