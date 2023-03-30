package main

import (
	"database/sql"
	"fmt"
	"formative-15/controllers"
	"formative-15/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Alwiassagafit46"
	dbname   = "practice"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	// ENV Configuration
	err := godotenv.Load("./config/.env")
	if err != nil {
		fmt.Println("failed to load environment file")
	} else {
		fmt.Println("successfully loaded environment file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run("localhost:8080")
}
