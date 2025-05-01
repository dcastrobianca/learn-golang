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
	conn, err := getConnection(dbUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to database.")

	// Query all albums
	albums, err := getAlbumByArtist(conn, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	// Print albums
	fmt.Println("Albums from John Coltrane:")
	for _, a := range albums {
		fmt.Printf("ID: %d | Title: %s | Artist: %s | Price: %.2f\n", a.ID, a.Title, a.Artist, a.Price)
	}

	// Hard-code ID 2 here to test the query.
	fmt.Println("Albums from id=4:")
	a, err := getAlbumByID(conn, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %d | Title: %s | Artist: %s | Price: %.2f\n", a.ID, a.Title, a.Artist, a.Price)

	// Add a new album
	newAlbum := Album{
		Title:  "From Zero",
		Artist: "Linkin Park",
		Price:  89.99,
	}
	if err := addAlbum(conn, newAlbum); err != nil {
		log.Fatal("failed to add album:", err)
	} else {
		fmt.Printf("Added new album")
	}
}

func getConnection(connStr string) (*pgx.Conn, error) {
	if connStr == "" {
		return nil, fmt.Errorf("connection string is empty")
	}
	// Connect to the database using pgx
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return conn, nil
}

func getAlbumByArtist(conn *pgx.Conn, artist string) ([]Album, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM album WHERE artist = $1", artist)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var a Album
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			return nil, fmt.Errorf("row scan failed: %v", err)
		}
		albums = append(albums, a)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading rows: %v", rows.Err())
	}

	return albums, nil
}

func getAlbumByID(conn *pgx.Conn, id int) (Album, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM album WHERE id = $1", id)
	if err != nil {
		return Album{}, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var a Album
	if rows.Next() {
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			return Album{}, fmt.Errorf("row scan failed: %v", err)
		}
	} else {
		return Album{}, fmt.Errorf("no album found with ID %d", id)
	}

	if rows.Err() != nil {
		return Album{}, fmt.Errorf("error reading rows: %v", rows.Err())
	}

	return a, nil
}

func addAlbum(conn *pgx.Conn, a Album) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", a.Title, a.Artist, a.Price)
	if err != nil {
		return fmt.Errorf("insert failed: %v", err)
	}
	return nil
}
