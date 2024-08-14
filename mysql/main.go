package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	// Connecta com a database MySQL, mais opcoes podem ser usadas por exemplo:
	// cfg := mysql.Config{
	// 	User:   os.Getenv("DBUSER"),
	// 	Passwd: os.Getenv("DBPASS"),
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "recordings",
	// }
	// `db` eh um objeto que representa a conexao com o banco de dados,
	// ele possui metodos para executar queries e transacoes.
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/coisas?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Primeira conexao com ping para ver se ta tudo ok.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Query para ser executada.
	query := "SELECT * FROM blahaj;"

	// Executa uma query e guarda o resultado em rows.
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	// Rows precisa ser fechado no final.
	defer rows.Close()

	// Itera sobre as linhas do resultado.
	for rows.Next() {
		var id int
		var name string
		var color string
		var size float64
		// Scan copia os valores das colunas para as variaveis passadas.
		err := rows.Scan(&id, &name, &color, &size)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID %d\tNome: %s\tCor: %s\tTamanho: %.2f\n", id, name, color, size)
	}

	// Verifica se houve algum erro durante o iteracao das linhas.
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Query de multiplas linhas pode usar uma struct e slices.
	// type Album struct {
	// 	ID     int64
	// 	Title  string
	// 	Artist string
	// 	Price  float32
	// }
	// var albums []Album
	// ...
	// Perceba que no codigo acima foi usado var id, var name, var color...
	// Agora os dados serao escritos em uma struct, como abaixo: &alb.ID, &alb.Title.
	// for rows.Next() {
	//     var alb Album
	//     if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
	//         return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	//     }
	// 	   albums = append(albums, alb)
	// }

	// Query de unica linha
	// var alb Album
	// `db.QueryRow` espera por uma query que retorne apenas uma row.
	// row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	// if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
	//     if err == sql.ErrNoRows {
	//         return alb, fmt.Errorf("albumsById %d: no such album", id)
	//     }
	//     return alb, fmt.Errorf("albumsById %d: %v", id, err)
	// }
	// return alb, nil
}
