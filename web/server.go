package server

import (
  "net/http"
  "html"
  "fmt"
  "log"
//  "strings"
)

type Server struct {
  port string
}

func (s* Server) Init(port string){
  s.port = port
}

func (s* Server) Run() {

  http.HandleFunc("/vis/", func(w http.ResponseWriter, r *http.Request) {
  	http.ServeFile(w, r, "web/" + r.URL.Path[1:])
  })
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })
  log.Fatal(http.ListenAndServe(":" + s.port, nil))
  fmt.Println("done")
}
