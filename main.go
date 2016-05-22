package main  //local webserver created using the build a wiki tutorial on golang.org
              //run using  go run <nameoffile>.go
              //then visit local host 8080

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, there", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
