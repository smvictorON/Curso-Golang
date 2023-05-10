package main

import (
	"database/sql"
	"fmt"
	"log"

	//import implicito usa o _
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	strConnection := "root:ilewml@/golangCRUD?charset=utf8&parseTime-True&loc=Local"
	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão está aberta!")

	lines, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()

	fmt.Println(lines)
}
