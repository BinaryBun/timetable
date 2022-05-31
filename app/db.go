package main

import ("fmt"
        "database/sql"

        _ "github.com/go-sql-driver/mysql")

func read_db (day, group string) [][5]string {
  pass := [][5]string {{"f", "f", "f", "f", "f"}}
  login, passwd, name := "root", "binarybun", "timetable"
  db, err := sql.Open("mysql", login+":"+passwd+"@tcp(127.0.0.1:3306)/"+name)
  defer db.Close()

  if err != nil {
    fmt.Println("Error: ", err.Error())
  } else {
    req := "select `time`, audit, `name`, prepod, `data` from `table` where "
    req += fmt.Sprintf("`day` = '%s' and `group` = '%s';", day, group)
    return requests(db, req)
  }
  return pass
}

func requests(db *sql.DB, req string) [][5]string {
  end_data := [][5]string {}
  var data [5]string

  incert, err := db.Query(req)
  defer incert.Close()

  if err != nil {
      fmt.Println("Error: ", err.Error())
  } else {
    for incert.Next() {
      nil := incert.Scan(&data[0], &data[1], &data[2], &data[3], &data[4])
      if nil == nil {}
      end_data = append(end_data, data)
    }
  }
  return end_data
}
