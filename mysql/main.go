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

	// Executa a query e guarda o resultado em rows.
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
}
