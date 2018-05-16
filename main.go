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

		resp := []string{}
		for result.Next() {
			line := UserLine{}
			err = result.Scan(&line.Course, &line.Account)
			assertNil(err)

			resp = append(resp, line.Course)
		}

		c.JSON(http.StatusOK, resp)
	}
}
