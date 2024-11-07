package main

import (
	"fmt"
    "strings"
	"github.com/a-h/templ"
	"megenisj/guessing-game/Handlers"
	"megenisj/guessing-game/State"
	"megenisj/guessing-game/Pages"
	"net/http"
    "net/url"
)

func main() {
	State.SetupState()
	//todo:// for development
	game := State.State{
		GameId:  1,
		InLobby: true,
		Turn:    1,
		Phase:   1,
		Guesser: "jimbo"}

	State.Insert(game)
	fmt.Println(game.Guesser)

	//http.Handle("/", templ.Handler(Pages.Index()))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decodedPath, err := url.PathUnescape(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid URL path", http.StatusBadRequest)
			return
		}
        fmt.Println(decodedPath)
		if strings.Contains(decodedPath,"/🚀🎮") {
			Handlers.StartGameHandler(w, r)
            return
		} else {
		    Pages.Index().Render(r.Context(), w)	
		}
	})

	http.Handle("/🎮", templ.Handler(Pages.Game(State.Select(1))))

	//starts the lobby - create rules / invite people
	http.HandleFunc("/🆕🎮", Handlers.NewGameHandler)

	//start game
	//http.HandleFunc("/🚀🎮", Handlers.StartGameHandler)

	//return a game (polled by game page)
	http.HandleFunc("/🔄🎮", Handlers.NewGameHandler)

	//handle clues
	http.HandleFunc("/🕵️💬", Handlers.ClueHandler)
	// handle guesses
	fmt.Println("listening on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
