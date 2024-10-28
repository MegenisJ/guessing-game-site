package Handlers

import (
    "fmt"
    "io"
    "net/http"
)

func ClueHandler(w http.ResponseWriter, r *http.Request){
    fmt.Printf("recived req ")
    io.WriteString(w, "Ok \n")
}
