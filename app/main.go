package main

import ("fmt"
        "net/http"
        "html/template")

type data struct {
  Group string;
  Mon [][]string;
}

func (d *data) set_data() {
  d.Group = "211-331"
  d.Mon = [][]string{
                      {"9:00-10:30",
                "https://online.mospolytech.ru/",
                "lms",
                "Дискретные структуры и компьютинг (Лекция)",
                "Будылина Евгения Александровна, Набебин Алексей Александрович"},
                      {"10:40-12:10",
                "https://yandex.ru/maps/213/moscow/?ll=37.646215%2C55.704464&mode=poi&poi%5Bpoint%5D=37.646083%2C55.704282&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D1737978898&z=20",
                "ав4607",
                "Организационное и правовое обеспечение информационной безопасности (Лекция)",
                "Пашина Алла Дмитриевна"}}
}

func home_page(w http.ResponseWriter, r *http.Request) {
  d := data{}
  d.set_data()
  t, _ := template.ParseFiles("templace/index.html")
  t.Execute(w, d)
}

func pageHeaders() {
  http.Handle("/styles/",
              http.StripPrefix("/styles/",
                               http.FileServer(http.Dir("./styles/"))))
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":8080", nil)
}

func main() {
  fmt.Println("Start")
  pageHeaders()
}
