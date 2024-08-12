package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	createTable()
}

func createTable() {
	sqlstr := `CREATE TABLE riverlevel(location varchar(5))`

	db, err := sql.Open("mysql", "root:GaBe6ear@tcp(localhost:3306)/test?charset=utf8")

	if err != nil {
		panic(err)
	}
	println("DB 結構已建立")

	_, err = db.Exec(sqlstr)

	if err != nil {
		panic(err)
	}
	println("資料表建立成功！")

	defer db.Close()
}
