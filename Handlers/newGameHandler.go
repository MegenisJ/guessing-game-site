package Handlers

import (
    "fmt"
    "net/http"
    "megenisj/guessing-game/Pages"
    "megenisj/guessing-game/State"
)

func NewGameHandler(w http.ResponseWriter, r *http.Request){
    fmt.Printf("recived req.... creating new game")

    if err:= r.ParseForm(); err !=nil {
        http.Error(w,"Unable to parse new game form", http.StatusBadRequest)
        return
    }
    //TODO: New game logic
    Pages.Game(State.New()).Render(r.Context(), w)
}
