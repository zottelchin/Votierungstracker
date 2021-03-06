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
	router.DELETE("/api/tasks/:account/:course", deleteCourse(db))

	router.Group("/api", func(c *gin.Context) {
		c.String(404, "API endpoint not found.\n")
	})
	router.Use(rewrite(router, `^/[a-zA-Z0-9-]+(?:/[^/]+)?$`, "/"))

	router.Run(":8900")
}

func deleteCourse(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		resultL, err := statements["delLimits"].Exec(c.Param("account"), c.Param("course"))
		fmt.Printf("Deleted for user %s limits for course %s \n", c.Param("account"), c.Param("course"))
		assertNil(err)

		resultC, err := statements["delCourse"].Exec(c.Param("account"), c.Param("course"))
		fmt.Printf("Deleted for user %s course %s\n", c.Param("account"), c.Param("course"))
		assertNil(err)

		numC, _ := resultC.RowsAffected()
		numL, _ := resultL.RowsAffected()
		c.JSON(http.StatusOK, gin.H{
			"Number deleted Rows from Courses": numC,
			"Number deleted Rows from Limits":  numL,
			"Deleted Course":                   c.Param("course"),
			"For User":                         c.Param("account"),
		})
	}
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
			err = result.Scan(&line.Course, &line.Account, &line.Points, &line.MaxPoints, &line.Min, &line.MinType)
			assertNil(err)

			resultLine := UserLineArr{}
			resultLine.Course = line.Course
			if line.MinType.Valid && line.MinType.String == "percent" {
				if line.MaxPoints > 0 {
					resultLine.AccPerc = float64(100*line.Points) / float64(line.MaxPoints)
				} else {
					resultLine.AccPerc = 0
				}
			}
			if line.MinType.Valid && line.MinType.String == "points" {
				resultLine.AccPerc = float64(line.Points)
			}
			resultLine.Min = line.Min.Float64
			resultLine.MinType = line.MinType.String
			if resultLine.MinType == "" {
				resultLine.MinType = "none"
			}

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

		var resp int
		for result.Next() {
			line := presStruct{}
			err = result.Scan(&line.Pres)
			assertNil(err)
			if line.Pres.Valid {
				resp = int(line.Pres.Int64)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"presentations": resp,
		})
	}
}

func getPerc(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := statements["getPerc"].Query(c.Param("course"), c.Param("account"))
		assertNil(err)
		defer result.Close()

		line := percStruct{}
		for result.Next() {
			err = result.Scan(&line.Minimum, &line.Type)
			assertNil(err)
		}

		c.JSON(http.StatusOK, line)
	}
}

func setPres(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		pres := presSetStruct{}
		err := c.BindJSON(&pres)
		assertNil(err)
		result, err := statements["setPres"].Exec(c.Param("course"), c.Param("account"), c.Param("course"), c.Param("account"), c.Param("course"), c.Param("account"), c.Param("course"), c.Param("account"), pres.Pres)
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
		result, err := statements["setPerc"].Exec(c.Param("course"), c.Param("account"), c.Param("course"), c.Param("account"), perc.Minimum, perc.Type, c.Param("course"), c.Param("account"))
		assertNil(err)

		id, _ := result.RowsAffected()

		c.JSON(http.StatusOK, gin.H{
			"updated percentage": id,
		})
	}
}
