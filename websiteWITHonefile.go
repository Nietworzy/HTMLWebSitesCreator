package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

    if r.RequestURI == "/" {
        fmt.Fprintf(w, `
        <html>
           <head>
              <title>My first site - Home</title>
              <meta charset="UTF-8">
              <rel name="stylesheet" href="mystyle.css">
           </head>
           <body>
              Hello on my first site!
           <body>
        </html>
    `)
   } else {
      fmt.Fprintf(w, `
      <html>
          <head>
              <title>My first site - Error</title>
          </head>
          <body>
              This site has does not exist!
          </body>
      </html>
    `)
   }
}

func main() {

    http.HandleFunc("/hello", hello)

    http.ListenAndServe(":2209", nil)
}
