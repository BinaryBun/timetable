package main

import ("fmt"
        "database/sql"

        _ "github.com/go-sql-driver/mysql")
type db struct {

}

func (d *db) read_db (day, group string) string  {

}

func main()  {
  fmt.Println("MySQL")

  db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)")
}
