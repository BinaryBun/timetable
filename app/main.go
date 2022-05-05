package main

import ("fmt"
        "net/http"
        "html/template")

type data struct {
  Group string;
}

func home_page(w http.ResponseWriter, r *http.Request) {
  d := data{"211-331"}
  t, _ := template.ParseFiles("templace/index.html")
  t.Execute(w, d)
}

func pageHeaders() {
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":8080", nil)
}

func main() {
  fmt.Println("Start")
  pageHeaders()
}
