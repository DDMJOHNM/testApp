package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	TEST "test/gen"
)

func main() {

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		GetAsString("DB_USER", "postgres"),
		GetAsString("DB_PASSWORD", "mysecretpassword"),
		GetAsString("DB_HOST", "localhost"),
		GetAsInt("DB_PORT", 5432),
		GetAsString("DB_NAME", "postgres"),
	)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping: ", err)
	}

	st := TEST.New(db)

	fmt.Println(st)
}
