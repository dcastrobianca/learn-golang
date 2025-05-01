package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}

func main() {
	dbUrl := "postgres://postgres:postgres@localhost:5332/recordings"
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected to database.")

	// Query all albums
	rows, err := conn.Query(context.Background(), "SELECT id, title, artist, price FROM album")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var a Album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			log.Printf("Row scan failed: %v\n", err)
			continue
		}
		albums = append(albums, a)
	}

	if rows.Err() != nil {
		log.Fatalf("Error reading rows: %v\n", rows.Err())
	}

	// Print albums
	fmt.Println("Albums:")
	for _, a := range albums {
		fmt.Printf("ID: %d | Title: %s | Artist: %s | Price: %.2f\n", a.ID, a.Title, a.Artist, a.Price)
	}
}
