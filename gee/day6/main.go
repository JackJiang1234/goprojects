package main

import (
	"example/gee"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		t := time.Now()
		ctx.Next()
		log.Printf("v2 logger [%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHtmlGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{
		Name: "jianyong",
		Age:  20,
	}
	stu2 := &student{
		Name: "jack",
		Age:  22,
	}
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/index", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>Index Page</h1>")
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.Html(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
