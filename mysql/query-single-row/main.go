package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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

	// Se o db nao for definido globalmente, ele sera passado como parametro
	alb, err := albumPorID(db, 23)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album encontrado: %v\n", alb)
}

func albumPorID(db *sql.DB, id int) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}

	return alb, nil
}
