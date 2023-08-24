package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"theprimeagen.tv/conway/pkg/database"
	"theprimeagen.tv/conway/pkg/pages"
)


type TemplateRenderer struct {
    templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

    funcMap := template.FuncMap{
        "loop": func(from, to int) <-chan int {
            ch := make(chan int)
            go func() {
                for i := from; i <= to; i++ {
                    ch <- i
                }
                close(ch)
            }()
            return ch
        },
    }

    tmpls, err := template.New("").Funcs(funcMap).ParseGlob("public/views/*.html")

    if err != nil {
        log.Fatalf("couldn't initialize templates: %v", err)
    }

    url := os.Getenv("DB_URL")
    if url == "" {
        url = "/tmp/conway.db"
    }

    err = database.InitDB(url)
    if err != nil {
        log.Fatalf("couldn't initialize db: %v", err)
    }

    e := echo.New()
    e.Renderer = &TemplateRenderer{
        templates: tmpls,
    }

    e.Use(middleware.Logger())
    e.Static("/dist", "dist");
    e.Static("/css", "css");

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index.html", nil)
    })

    e.GET("/saved", pages.Saved)

    e.POST("/save", func(c echo.Context) error {
        return c.Render(200, "saved-msg", nil)
    })

    e.POST("/saveAt", func(c echo.Context) error {
        return c.Render(200, "saved-msg", nil)
    })


    e.Logger.Fatal(e.Start(":42069"))
}

