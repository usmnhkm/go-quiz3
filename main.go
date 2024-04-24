package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz3/database"
	"quiz3/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed load file env")
	} else {
		fmt.Println("Success load file env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	fmt.Println(psqlInfo)
	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}
	defer DB.Close()

	database.DbMigrate(DB)
	fmt.Println("App listening on port : " + os.Getenv("PORT"))
	routes.StartServer().Run(":" + os.Getenv("PORT"))
}
