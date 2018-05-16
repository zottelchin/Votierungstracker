package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var statements = map[string]*sql.Stmt{}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil || db == nil {
		panic(err)
	}

	statements["getTasks"] = mustPrepare(db, "SELECT id, name, points, maxPoints, presentations FROM tasks WHERE user = ? AND class = ?")
	statements["addTask"] = mustPrepare(db, "INSERT INTO tasks(user, class, name, points, maxPoints, presentations) VALUES(?, ?, ?, ?, ?, ?)")
	statements["updateTask"] = mustPrepare(db, "UPDATE tasks SET name = ?, points = ?, maxPoints = ?, presentations = ? WHERE user = ? AND class = ? AND id = ?")
	statements["deleteTask"] = mustPrepare(db, "DELETE FROM tasks WHERE user = ? AND class = ? AND id = ?")
	statements["getCourses"] = mustPrepare(db, "SELECT class, user FROM tasks WHERE user = ? GROUP BY class")

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		class VARCHAR NOT NULL,
		user VARCHAR NOT NULL,
		name VARCHAR NOT NULL,
		points INTEGER NOT NULL,
		maxPoints INTEGER NOT NULL,
		presentations INTEGER NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

func mustPrepare(db *sql.DB, sql string) *sql.Stmt {
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	return stmt
}
