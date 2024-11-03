package Handlers

import (
    "fmt"
    "io"
    "net/http"
)

func NewGameHandler(w http.ResponseWriter, r *http.Request){
    fmt.Printf("recived req.... creating new game")
    if err:= r.ParseForm(); err !=nil {
        http.Error(w,"Unable to parse new game form", http.StatusBadRequest)
        return
    }
    io.WriteString(w, "Ok \n")
}
