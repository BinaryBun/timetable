package main

import ("fmt"
        "database/sql"

        _ "github.com/go-sql-driver/mysql")
/*type db struct {}

func (d *db) read_db (day, group string) string  {
  return "pass"
}*/

func requests(db *sql.DB, req string) {
  var name, passwd string

  incert, err := db.Query(req)
  defer incert.Close()

  if err != nil {
      fmt.Println("Error: ", err.Error())
  } else {
    for incert.Next() {
      nil := incert.Scan(&name, &passwd)
      if nil == nil {}
      fmt.Println(name, passwd)
    }
  }
}

func main() {
  login, passwd := "root", "binarybun"
  db, err := sql.Open("mysql", login+":"+passwd+"@tcp(127.0.0.1:3306)/parser")
  //fmt.Println(fmt.Sprintf("%T\n%T", db, err))

  if err != nil {
    fmt.Println("Error: ", err.Error())
  } else {
    requests(db, "SELECT * from passwdords;")
  }
  defer db.Close()
}
