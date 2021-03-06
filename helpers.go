package main

import (
	"database/sql"
	"regexp"

	"github.com/gin-gonic/gin"
)

type GridLine struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Points        int    `json:"points"`
	MaxPoints     int    `json:"maxPoints"`
	Presentations int    `json:"presentations"`
}

type UserLine struct {
	Course    string          `json:"course"`
	Account   string          `json:"account"`
	Points    int             `json:"Points"`
	MaxPoints int             `json:"maxPoints"`
	Min       sql.NullFloat64 `json:"minimum"`
	MinType   sql.NullString  `json:"minimumType"`
}

type UserLineArr struct {
	Course  string  `json:"course"`
	AccPerc float64 `json:"actual"`
	Min     float64 `json:"minimum"`
	MinType string  `json:"minimumType"`
}

type percStruct struct {
	Type    string  `json:"type"`
	Minimum float64 `json:"minimum"`
}

type presStruct struct {
	Pres sql.NullInt64 `json:"presentations"`
}

type presSetStruct struct {
	Pres int `json:"presentations"`
}

func assertNil(x interface{}) {
	if x != nil {
		panic(x)
	}
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
