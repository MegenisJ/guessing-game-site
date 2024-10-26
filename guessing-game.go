package main

import (
	"fmt"
	"net/http"
	"github.com/a-h/templ"
	pages "megenisj/guessing-game/pages"
)

func main() {
    setupState();
    //index
    index:= pages.Index()
    http.Handle("/", templ.Handler(index))

    //New game page - create rules / invite people 
    new:= pages.New()
    http.Handle("/new", templ.Handler(new))

    //New game page - create rules / invite people 
    game:= pages.Game()
    http.Handle("/game", templ.Handler(game))
    fmt.Println("listening on http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
