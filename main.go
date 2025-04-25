package main

import (
	"api-simgos/database"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	db  *sql.DB
	db2 *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	// connect into postgresql database
	db, err = sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection to database is failed")
		panic(err)
	}

	fmt.Println("Successfully make connection to database")

	satuInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("SATUHOST"),
		os.Getenv("SATUPORT"),
		os.Getenv("SATUUSER"),
		os.Getenv("SATUPASSWORD"),
		os.Getenv("SATUDATABASE"))

	// connect into postgresql database
	db2, err = sql.Open("postgres", satuInfo)
	err = db2.Ping()
	if err != nil {
		fmt.Println("Connection to database satu is failed")
		panic(err)
	}

	fmt.Println("Successfully make connection to database satu")

	// database migration
	database.DbMigrate(db)
	database.DbMigrateSatu(db2)
	defer db.Close()
	defer db2.Close()

	// start routers from routers.go
	StartServer(":" + os.Getenv("PORT"))
}
