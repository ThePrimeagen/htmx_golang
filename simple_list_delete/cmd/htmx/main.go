package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"theprimeagen.tv/go_htmx/pkg/view"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Item struct {
    Name string
    Count int
    Id int
}

type Content struct {
    Items []Item
}

func find_item(items []Item, id int) *Item {
    for i := range items {
        if items[i].Id == id {
            return &items[i]
        }
    }
    return nil
}

func delete_item(items []Item, id int) []Item {
    for i := range items {
        if items[i].Id == id {
            return append(items[:i], items[i+1:]...)
        }
    }
    return items
}

func main() {
	e := echo.New()

    tmpl := template.New("index")
	var err error
	if tmpl, err = tmpl.Parse(view.Index); err != nil {
		fmt.Println(err)
	}

	if tmpl, err = tmpl.Parse(view.Items); err != nil {
		fmt.Println(err)
	}

	if tmpl, err = tmpl.Parse(view.Item); err != nil {
		fmt.Println(err)
	}

	if tmpl, err = tmpl.Parse(view.ItemCount); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		e.StdLogger.Fatal(err)
	}

    e.Use(middleware.Logger())

	e.Renderer = &TemplateRenderer{
		templates: tmpl,
	}

    items := Content{
        Items: []Item{
            {Name: "Item 1", Count: 1, Id: 1},
            {Name: "Item 2", Count: 2, Id: 2},
            {Name: "Item 3", Count: 3, Id: 3},
        },
    }

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", items)
	})

    e.DELETE("/item/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return err
        }

        item := find_item(items.Items, id)
        if item == nil {
            return echo.NewHTTPError(http.StatusNotFound)
        }

        items.Items = delete_item(items.Items, id)
        return c.NoContent(http.StatusOK)
    });

    e.POST("/count/:id", func(c echo.Context) error {

        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return err
        }

        item := find_item(items.Items, id)
        item.Count += 1

		return c.Render(http.StatusOK, "item-count", item)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
