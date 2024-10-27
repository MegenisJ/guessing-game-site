package main

import (
	"fmt"
	handlers "megenisj/guessing-game/Handlers"
	state "megenisj/guessing-game/State"
	pages "megenisj/guessing-game/pages"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
    state.SetupState();
    //index
    game := state.State{
        GameId: 1,
        Turn: 1,
        Phase: 1,
        Guesser: "jimbo"}

    state.Insert(game)

    fmt.Println(game.Guesser)

    http.Handle("/", templ.Handler(pages.Index()))

    //New game page - create rules / invite people 
    http.Handle("/new", templ.Handler(pages.New()))

    //New game page - create rules / invite people 
    http.Handle("/game", templ.Handler(pages.Game(game)))

    http.HandleFunc("/clue", handlers.ClueHandler) 

    fmt.Println("listening on http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
