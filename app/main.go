package main

import ("fmt"
        "net/http"
        "html/template")

func home_page(w http.ResponseWriter, r *http.Request) {
  t, nil := template.ParseFiles("templace/index.html")
  t.Execute(w, nil)
}

func pageHeaders() {
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":8080", nil)
}

func main() {
  fmt.Println("Start")
  pageHeaders()
}
