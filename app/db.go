package main

import ("fmt"
        "database/sql"

        _ "github.com/go-sql-driver/mysql")
/*type db struct {}

func (d *db) read_db (day, group string) string  {
  return "pass"
}*/

func requests(db *sql.DB) {

}

func main() {
  db, err := sql.Open("mysql", "root:binarybun@tcp(127.0.0.1:3306)/parser")
  ///fmt.Println(fmt.Sprintf("%T\n%T", db, err))

  if err != nil {
      fmt.Println("Error: ", err.Error())
  } else {
    requests(db)
  }

  defer db.Close()
}
