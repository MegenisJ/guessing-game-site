
package Handlers

import (
    "fmt"
    "net/http"
    "megenisj/guessing-game/Pages"
    "megenisj/guessing-game/State"
    "strconv"
)

func StartGameHandler(w http.ResponseWriter, r *http.Request){
    fmt.Printf("recived req.... starting new game")
    gameId := r.URL.Query().Get("game") 
    fmt.Println(gameId)
    i, err := strconv.Atoi(gameId)
    if err == nil {
        http.Error(w,"Unable to parse new game form. Must be int", http.StatusBadRequest)
        return
    }
    State.Start(i)
    Pages.Game(State.Select(i)).Render(r.Context(), w)
}
