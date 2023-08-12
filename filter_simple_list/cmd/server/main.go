package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
    _ "github.com/libsql/libsql-client-go/libsql"
)

type Item struct {
    Name string
    Id int
    Count int
}

type Content struct {
    Items []Item
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var ids = 0

func main() {
    tmpl, err := template.ParseGlob("./public/views/*.html")

    if err != nil {
        log.Fatalf("unable to parse glob %e\n", err)
    }

    e := echo.New()
    e.Renderer = &TemplateRenderer{
        templates: tmpl,
    }

    items := Content{
        Items: []Item{},
    }

    e.POST("/items", func(c echo.Context) error {
        name := c.FormValue("name")
        if name == "" {
            return c.String(http.StatusBadRequest, "name is empty")
        }

        item := Item{
            Name: name,
            Id: ids,
            Count: 0,
        }
        items.Items = append(items.Items, item)

        ids += 1

        return c.Render(http.StatusOK, "item.html", item)
    })

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index.html", items)
    })

    e.Logger.Fatal(e.Start(":42069"))
}

