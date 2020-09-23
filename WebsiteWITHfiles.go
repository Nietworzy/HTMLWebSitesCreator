package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

    if r.RequestURI == "/" {
        http.ServeFile(w, r, "index.html")
    } else {
        http.ServeFile(w, r, "error404.html")
    }
}

func main() {

    http.HandleFunc("/hello", hello)

    http.ListenAndServe(":2209", nil)
}
