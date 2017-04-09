package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {

  rtr := mux.NewRouter()
  rtr.Handle("/", rootHandler()).Methods("GET")

  http.Handle("/", rtr)

  if err := http.ListenAndServe(":8080", Log(http.DefaultServeMux)); err != nil {
      log.Fatal("ListenAndServe: ", err)
  }
}

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
          handler.ServeHTTP(w, r)
    })
}

func rootHandler() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, New World!"))
  })
}
