package main

import (
  "log"
  "net/http"
  "runtime"
  "path"
//  "os"
  "github.com/gorilla/mux"
  "compress/gzip"
  "html/template"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/mattn/go-sqlite3"
//  "github.com/stretchr/graceful"
  "github.com/carbocation/interpose"
  "github.com/carbocation/interpose/middleware"
//  "code.google.com/p/gorilla/sessions"
  "github.com/gorilla/sessions"
)
var db, drr = gorm.Open("postgres", "user=railerde dbname=aleph0 sslmode=disable password=liberatorarchon7!4$K")

//var f, lrr = os.OpenFile("logfile.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
//var defer f.Close()

//var log.SetOutput(f)
//var store sessions.Store
var store = sessions.NewCookieStore([]byte("yomamareallyuglynoscopehomedawgyoumadbrauwotm8lolhahahasanyonereallygonefarevenasdecidedtouseevengowanttodolookmorelike"))
func main(){
  runtime.GOMAXPROCS(runtime.NumCPU())
//  ApikeyCreate("admin")
  router := NewRouter()
//  router.PathPrefix("/public").Handler(http.FileServer(http.Dir("./public/")))
  dbsetup()
  middle := interpose.New()
  middle.Use(middleware.GorillaLog())
  middle.Use(middleware.NegroniGzip(gzip.DefaultCompression))
  littlemiddle := interpose.New()
    littlemiddle.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
            rw.Header().Set("X-Server-Name", "א --Dappit !wut-wut! Server-- א")
            //put middleware stuff here
            apikeys := Apikeys{}
            token := req.FormValue("token")
            db.Where("secret = ?", token).Find(&apikeys)
            //
           if len(apikeys) != 0 {
             next.ServeHTTP(rw, req)
           }
           if len(apikeys) == 0 {
             http.Error(rw, "401 Unauthorized", 401)
           }
        })
    })
  littlemiddle.UseHandler(router)
//  UserCreate("username", "password")
  routd := mux.NewRouter()
  pux := PewRouter()
  routd.PathPrefix("/api/v1").Handler(littlemiddle)
  routd.PathPrefix("/api/voo").Handler(pux)
  routd.HandleFunc("/", redirect) 
  routd.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/"))) 
//  routd.PathPrefix("/").Handler(HandleFunc("/", serveTemplate))
//  http.HandleFunc("/", serveTemplate)
  middle.UseHandler(routd)
  log.Println("Listening...")
//  go log.Fatal(http.ListenAndServe(":8080", middle))
  go http.ListenAndServe(":8080", middle)
  go log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", middle))
//  go log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", middle))
//    graceful.Run(":8080", 10*time.Second, middle)
}

func redirect(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/index.htm", 301)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("templates", "layout.html")
  fp := path.Join("templates", r.URL.Path)

  tmpl, _ := template.ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}
