package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/mattn/go-sqlite3"
)

func setupState(){
    os.Remove("./foo.db")
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    sqlStmt := `
    create table foo (id integer not null primary key, state integer);
    delete from foo;
    `
    _, err = db.Exec(sqlStmt)

    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }

    stmt, err := tx.Prepare("insert into foo(id, state) values(?, ?)")

    if err != nil {
        log.Fatal(err)
    }

    defer stmt.Close()
    _,err = stmt.Exec(1,10)
    if err != nil {
        log.Fatal(err)
    }

    err = tx.Commit()
    if err != nil{
        log.Fatal(err)
    }

    rows,err := db.Query("select id,state from foo")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next(){
        var id int
        var state int
        err = rows.Scan(&id, &state)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Println(id,state)
    }
}
