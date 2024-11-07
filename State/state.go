package State

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type State struct {
	GameId  int
	InLobby bool
	Turn    int
	Phase   int
	Guesser string
}

func SetupState() {
	os.Remove("./foo.db")
	db := connect()
	defer db.Close()

	sqlStmt := `
        create table foo (GameId integer not null primary key, InLobby bool, Phase integer, TurnCount int, Guesser string);
        delete from foo;
    `
	_, err := db.Exec(sqlStmt)

	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func New() State {
	game := State{
		GameId:  1,
		InLobby: true,
		Turn:    1,
		Phase:   1,
		Guesser: "jimbo"}

	Insert(game)
	return game
}

func Start(game int) {

	db := connect()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("update foo set InLobby = false where GameId = 1")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(s State) {

	db := connect()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(InLobby, Phase, TurnCount, Guesser) values(?,?, ?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.InLobby,s.Phase, s.Turn, s.Guesser)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func Select(gameId int) State {
	db := connect()
	defer db.Close()

	rows, err := db.Query("select GameId,Phase, TurnCount,Guesser from foo where GameId == 1")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var GameId int
		var Phase int
		var TurnCount int
		var Guesser string
		err = rows.Scan(&GameId, &Phase, &TurnCount, &Guesser)
		if err != nil {
			log.Fatal(err)
		}
		return State{GameId: GameId, Phase: Phase, Turn: TurnCount, Guesser: Guesser}
	}
	panic("unable to find record")
}

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}

	return db
}
