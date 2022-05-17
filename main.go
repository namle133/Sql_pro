package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=SQL_PROJECT user=postgres password=Namle311")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v", err))
	}
	defer conn.Close()
	log.Println("Connected to database")
}
