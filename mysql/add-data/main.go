package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/coisas?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	newAlb, err := addAlbum(db, &Album{Title: "Titulo", Artist: "Artista", Price: 49.99})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID do album adcionado: %v\n", newAlb)
}

func addAlbum(db *sql.DB, alb *Album) (id int64, err error) {
	// Sempre usar `?` em comandos para evitar SQL Injection.
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	// LastInsertId retorna o ID do ultimo registro inserido.
	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}
