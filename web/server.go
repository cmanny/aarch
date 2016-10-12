package server

import (
  "net/http"
  "fmt"
  "log"
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

  http.HandleFunc("/api/", func(w, http.ResponseWriter, r* http.Request) {
    /* Fill in with json requests */
    return
  })

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "web/index.html")
  })
  log.Fatal(http.ListenAndServe(":" + s.port, nil))
  fmt.Println("done")
}
