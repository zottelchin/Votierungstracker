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

	migrate(db)

	statements["getTasks"] = mustPrepare(db, "SELECT id, name, points, maxPoints, presentations FROM tasks WHERE user = ? AND class = ?")
	statements["addTask"] = mustPrepare(db, "INSERT INTO tasks(user, class, name, points, maxPoints, presentations) VALUES(?, ?, ?, ?, ?, ?)")
	statements["updateTask"] = mustPrepare(db, "UPDATE tasks SET name = ?, points = ?, maxPoints = ?, presentations = ? WHERE user = ? AND class = ? AND id = ?")
	statements["deleteTask"] = mustPrepare(db, "DELETE FROM tasks WHERE user = ? AND class = ? AND id = ?")
	statements["getCourses"] = mustPrepare(db, "SELECT tasks.class, tasks.user, SUM(tasks.points), SUM(tasks.maxPoints), limits.percent FROM tasks left JOIN limits ON tasks.user = limits.user and tasks.class = limits.class WHERE tasks.user = ? GROUP BY tasks.class")
	statements["setPerc"] = mustPrepare(db, "REPLACE INTO limits (id, class, user, percent, presentations) VALUES ((SELECT id FROM limits WHERE class = ? and user = ?), ?, ?, ?, (SELECT presentations FROM limits WHERE class = ? and user = ?))")
	statements["getPerc"] = mustPrepare(db, "SELECT percent FROM limits WHERE class = ? and user = ?")
	statements["setPres"] = mustPrepare(db, "REPLACE INTO limits (id, class, user, percent, presentations) VALUES ((SELECT id FROM limits WHERE class = ? and user = ?), ?, ?, (SELECT percent FROM limits WHERE class = ? and user = ?), ?)")
  statements["getPres"] = mustPrepare(db, "SELECT presentations FROM limits WHERE class = ? and user = ?")
  statements["delCourse"] = mustPrepare(db, "DELETE FROM tasks WHERE user = ? AND class = ?")
  statements["delLimits"] = mustPrepare(db, "DELETE FROM limits WHERE user = ? AND class = ?")

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
    );`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}

	sql2 := `
	CREATE TABLE IF NOT EXISTS limits(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		class VARCHAR NOT NULL,
		user VARCHAR NOT NULL,
		percent INTEGER,
		presentations INTEGER 
	);`

	_, err2 := db.Exec(sql2)
	if err2 != nil {
		panic(err2)
	}
}

func mustPrepare(db *sql.DB, sql string) *sql.Stmt {
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	return stmt
}
