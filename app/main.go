package main

import ("log"
        "fmt"
        "net/http"
        "html/template"
        "time")

type data struct {
  Group string;
  Mon [][5]string;
  Tue [][5]string;
  Wed [][5]string;
  Thu [][5]string;
  Fri [][5]string;
  Sat [][5]string;
}

type tmp struct {
  Ctn int
}

var Ctn = 0
const group = "211-331"

func (d *data) set_data() {
  d.Mon = read_db("Понедельник", d.Group)
  d.Tue = read_db("Вторник", d.Group)
  d.Wed = read_db("Среда", d.Group)
  d.Thu = read_db("Четверг", d.Group)
  d.Fri = read_db("Пятница", d.Group)
  d.Sat = read_db("Суббота", d.Group)
}

func home_page(w http.ResponseWriter, r *http.Request) {
  //fmt.Println(r.RemoteAddr)
  //read cookie
  group_cook, err := r.Cookie("group")
  value := ""
  if err != nil {
		log.Println("Error occured while reading cookie")
    value = group
	} else {
    value = group_cook.Value
  }
  //set data
  d := data{}
  d.Group = value
  d.set_data()
  t, _ := template.ParseFiles("templace/index.html")
  t.Execute(w, d)
  Ctn ++
}

func count_page(w http.ResponseWriter, r *http.Request)  {
  t, _ := template.ParseFiles("templace/ctn.html")
  data := tmp {Ctn/2}
  t.Execute(w, data)
  Ctn --
}

func press(w http.ResponseWriter, r *http.Request)  {
  // log
  log.Println(">> get data")
  log.Println(r.FormValue("group_inp"))
  // set coolie
  http.SetCookie(w, &http.Cookie{
			Name:       "group",
			Value:      r.FormValue("group_inp"),
      Path:       "/",
  		Domain:     "",
  		Expires:    time.Time{},
  		RawExpires: "",
  		MaxAge:     0,
      Secure:     false,
      HttpOnly:   true,
  		SameSite:   0,
  		Raw:        "",
  		Unparsed:   nil,})
  //redirect
  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func pageHeaders() {
  http.Handle("/styles/",
              http.StripPrefix("/styles/",
                               http.FileServer(http.Dir("./styles/"))))
  http.HandleFunc("/", home_page)
  http.HandleFunc("/ctn/", count_page)
  http.HandleFunc("/press/", press)
  http.ListenAndServe(":8080", nil)
}

func main() {
  fmt.Println("Start")
  pageHeaders()
}
