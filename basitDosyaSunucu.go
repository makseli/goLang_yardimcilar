package main

import "net/http"

func main() {
        http.ListenAndServe(":81", http.FileServer(http.Dir(".")))
}
