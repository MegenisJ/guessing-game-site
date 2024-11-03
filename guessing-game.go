package main

import (
	"fmt"
	"net/http"
	"github.com/a-h/templ"
	handlers "megenisj/guessing-game/Handlers"
	state "megenisj/guessing-game/State"
	pages "megenisj/guessing-game/pages"
)

func main() {
    state.SetupState();
    game := state.State{
        GameId: 1,
        InLobby: true,
        Turn: 1,
        Phase: 1,
        Guesser: "jimbo"}

    state.Insert(game)
    fmt.Println(game.Guesser)

    http.Handle("/", templ.Handler(pages.Index()))

    //New game page - create rules / invite people 
    http.Handle("/new", templ.Handler(pages.New(game)))
//    http.HandleFunc("/new", handlers.NewGameHandler)  

    http.HandleFunc("/start", handlers.NewGameHandler) 

    //New game page - create rules / invite people 
    http.Handle("/game", templ.Handler(pages.Game(state.Select(1))))

    http.HandleFunc("/clue", handlers.ClueHandler) 

    fmt.Println("listening on http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
