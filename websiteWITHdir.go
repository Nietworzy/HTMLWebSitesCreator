package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("."))
  http.Handle("/", fs)

  log.Println("Listening on :2209...")
  err := http.ListenAndServe(":2209", nil)
  if err != nil {
    log.Fatal(err)
  }
}
