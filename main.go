package main

import (
	"database/sql"
	"net/http"
	"regexp"

	"github.com/gin-contrib/cors"
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

type UserData struct {
	UserLines []UserLine `json:"items"`
}

type UserLine struct {
	CLASS string `json:"class"`
	USER  string `json:"user"`
}

func main() {

	db := initDB("db/storage.db")
	migrate(db)
	router := gin.Default()

	router.StaticFile("/", "frontend/dist/index.html")
	router.Static("/static/", "frontend/dist/static/")

	router.Use(cors.Default())

	router.GET("/api/getVotes/:user/:class", func(c *gin.Context) {
		getUserTable(c, db)
	})
	router.GET("/api/deleteVotes/:id", func(c *gin.Context) {
		deleteRow(c, db)
	})
	router.GET("/api/addVotes/:user/:class/:task/:points/:max/:pres", func(c *gin.Context) {
		addRow(c, db)
	})
	router.GET("/api/getUser/:user", func(c *gin.Context) {
		getUser(c, db)
	})

	router.Use(rewrite(router, `^/[a-zA-Z0-9-]+(?:/[^/]+)?`, "/"))

	router.Run(":8900")
}

func rewrite(engine *gin.Engine, match string, target string) gin.HandlerFunc {
	expr := regexp.MustCompile(match)
	return func(c *gin.Context) {
		if expr.MatchString(c.Request.URL.Path) {
			c.Request.URL.Path = target
			engine.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
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

func getUser(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	sql := "SELECT class, user FROM tasks WHERE user = '" + user + "' GROUP BY class;"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	resp := UserData{}
	for rows.Next() {
		line := UserLine{}
		err2 := rows.Scan(&line.CLASS, &line.USER)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		resp.UserLines = append(resp.UserLines, line)
	}

	c.JSON(http.StatusOK, resp)
}
