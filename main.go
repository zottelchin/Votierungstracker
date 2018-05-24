package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	db := initDB("db/storage.db")
	migrate(db)
	router := gin.Default()

	if gin.Mode() != "release" {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"http://localhost:8999", "http://127.0.0.1:8999", "http://[::]:8999", "http://[::1]:8999"}
		config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
		router.Use(cors.New(config))
	}

	router.StaticFile("/", "frontend/dist/index.html")
	router.Static("/static/", "frontend/dist/static/")

	router.GET("/api/tasks/:account/:course", getTasks(db))
	router.POST("/api/tasks/:account/:course", addTask(db))
	router.PUT("/api/tasks/:account/:course/:id", updateTask(db))
	router.DELETE("/api/tasks/:account/:course/:id", deleteTask(db))
	router.GET("/api/courses/:account", getCourses(db))
	router.GET("/api/pres/:account/:course", getPres(db))
	router.PUT("/api/pres/:account/:course", setPres(db))
	router.GET("/api/perc/:account/:course", getPerc(db))
	router.PUT("/api/perc/:account/:course", setPerc(db))

	router.Group("/api", func(c *gin.Context) {
		c.String(404, "API endpoint not found.\n")
	})
	router.Use(rewrite(router, `^/[a-zA-Z0-9-]+(?:/[^/]+)?$`, "/"))

	router.Run(":8900")
}

func getTasks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := statements["getTasks"].Query(c.Param("account"), c.Param("course"))
		fmt.Printf("%+v %s %s\n", result, c.Param("account"), c.Param("course"))
		assertNil(err)
		defer result.Close()

		resp := []GridLine{}
		for result.Next() {
			line := GridLine{}
			err = result.Scan(&line.Id, &line.Name, &line.Points, &line.MaxPoints, &line.Presentations)
			assertNil(err)
			resp = append(resp, line)
		}

		c.JSON(http.StatusOK, resp)
	}
}

func addTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		task := GridLine{}
		err := c.BindJSON(&task)
		fmt.Printf("%+v %s %s\n", task, c.Param("account"), c.Param("course"))
		assertNil(err)

		result, err := statements["addTask"].Exec(c.Param("account"), c.Param("course"), task.Name, task.Points, task.MaxPoints, task.Presentations)
		assertNil(err)

		id, _ := result.LastInsertId()
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func updateTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		task := GridLine{}
		err := c.BindJSON(&task)
		assertNil(err)

		result, err := statements["updateTask"].Exec(task.Name, task.Points, task.MaxPoints, task.Presentations, c.Param("account"), c.Param("course"), c.Param("id"))
		assertNil(err)

		num, _ := result.RowsAffected()
		c.JSON(http.StatusOK, gin.H{
			"updated": num,
		})
	}
}

func deleteTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := statements["deleteTask"].Exec(c.Param("account"), c.Param("course"), c.Param("id"))
		assertNil(err)

		num, _ := result.RowsAffected()
		c.JSON(http.StatusOK, gin.H{
			"deleted": num,
		})
	}
}

func getCourses(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := statements["getCourses"].Query(c.Param("account"))
		assertNil(err)
		defer result.Close()

		resp := []UserLineArr{}
		for result.Next() {
			line := UserLine{}
			err = result.Scan(&line.Course, &line.Account, &line.Points, &line.MaxPoints, &line.Per)
			assertNil(err)

			resultLine := UserLineArr{}
			resultLine.Course = line.Course
			resultLine.AccPerc = 100 * line.Points / line.MaxPoints
			resultLine.Per = int(line.Per.Int64)
			resultLine.Show = line.Per.Valid

			resp = append(resp, resultLine)
		}

		c.JSON(http.StatusOK, resp)
	}
}

func getPres(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := statements["getPres"].Query(c.Param("course"), c.Param("account"))
		assertNil(err)
		defer result.Close()

		resp := []int{}
		for result.Next() {
			line := presStruct{}
			err = result.Scan(&line.Pres)
			assertNil(err)
			resp = append(resp, line.Pres)
		}

		c.JSON(http.StatusOK, gin.H{
			"presentations": resp[0],
		})
	}
}

func getPerc(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := statements["getPerc"].Query(c.Param("course"), c.Param("account"))
		assertNil(err)
		defer result.Close()

		resp := []int{}
		for result.Next() {
			line := percStruct{}
			err = result.Scan(&line.Perc)
			assertNil(err)
			resp = append(resp, line.Perc)
		}

		c.JSON(http.StatusOK, gin.H{
			"percentage": resp[0],
		})
	}
}

func setPres(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		pres := presStruct{}
		err := c.BindJSON(&pres)
		assertNil(err)
		result, err := statements["setPres"].Exec(c.Param("course"), c.Param("account"), c.Param("course"), c.Param("account"), c.Param("course"), c.Param("account"), pres.Pres)
		assertNil(err)

		id, _ := result.RowsAffected()

		c.JSON(http.StatusOK, gin.H{
			"updated presentations": id,
		})
	}
}

func setPerc(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		perc := percStruct{}
		err := c.BindJSON(&perc)
		assertNil(err)
		result, err := statements["setPerc"].Exec(c.Param("course"), c.Param("account"), c.Param("course"), c.Param("account"), perc.Perc, c.Param("course"), c.Param("account"))
		assertNil(err)

		id, _ := result.RowsAffected()

		c.JSON(http.StatusOK, gin.H{
			"updated percentage": id,
		})
	}
}
