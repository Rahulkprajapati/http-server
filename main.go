package main

import (
    "net/http"
    "fmt"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}


func main() {
    http.HandleFunc("/", helloWorldPage)

}