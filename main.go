package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type GridLine struct {
	ID            int    `json:"id"`
	TASK          string `json:"task"`
	POINTS        int    `json:"points"`
	MAXPOINTS     int    `json:"maxPoints"`
	PRESENTATiONS int    `json:"pres"`
}

type DataCollection struct {
	GridLines []GridLine `json:"items"`
}

func main() {

	db := initDB("storage.db")
	migrate(db)
	router := gin.Default()

	router.StaticFile("/class", "static/index.html")
	router.StaticFile("/vue.js", "static/vue.js")
	router.StaticFile("/style.css", "static/style.css")
	router.StaticFile("/fontawesome-all.css", "static/css/fontawesome-all.css")
	router.Static("/webfonts", "static/webfonts/")

	router.GET("/api/getVotes/:user/:class", func(c *gin.Context) {
		getUserTable(c, db)
	})
	router.GET("/api/deleteVotes/:id", func(c *gin.Context) {
		deleteRow(c, db)
	})
	router.GET("api/addVotes/:user/:class/:task/:points/:max/:pres", func(c *gin.Context) {
		addRow(c, db)
	})

	router.Run(":8900")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
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

func getUserTable(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	class := c.Param("class")

	sql := "SELECT id, name, points, maxPoints, presentations FROM tasks WHERE class = '" + class + "' and user = '" + user + "';"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	resp := DataCollection{}
	for rows.Next() {
		line := GridLine{}
		err2 := rows.Scan(&line.ID, &line.TASK, &line.POINTS, &line.MAXPOINTS, &line.PRESENTATiONS)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		resp.GridLines = append(resp.GridLines, line)
	}

	c.JSON(http.StatusOK, resp)
}

func addRow(c *gin.Context, db *sql.DB) {
	sql := "INSERT INTO tasks(user, class, name, points, maxPoints, presentations) VALUES(?, ?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(c.Param("user"), c.Param("class"), c.Param("task"), c.Param("points"), c.Param("max"), c.Param("pres"))
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"created": id,
		})
	}

}

func deleteRow(c *gin.Context, db *sql.DB) {
	sql := "DELETE FROM tasks WHERE id = (?) "

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Param("id"))
	if err != nil {
		panic(err)
	}
	if err == nil {
		c.String(http.StatusOK, "Deleted")
	}
}
