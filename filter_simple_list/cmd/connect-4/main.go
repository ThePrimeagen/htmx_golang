package main

import (
	"html/template"
	"log"
	"os"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"theprimeagen.tv/filter_simple_list/pkg/database"
	"theprimeagen.tv/filter_simple_list/pkg/endpoints"
)

type User struct {
    name string
    wins int
}

type Page struct {
    user *User
}

func main() {
    db_url := os.Getenv("DB_URL")
    if db_url == "" {
        db_url = "file:///tmp/connect-4"
    }

    err := database.InitDB(db_url)
    if err != nil {
        log.Fatalf("could not initialize db: %+v", err)
    }

    tmpl, err := template.ParseFiles(
        "./public/views/connect-4/index.html",
        "./public/views/connect-4/name.html",
    )

    if err != nil {
        log.Fatalf("could not initialize templates: %+v", err)
    }

    e := echo.New()
    e.Renderer = endpoints.NewTemplateRenderer(tmpl)
    e.Use(middleware.Logger())

    e.GET("/", endpoints.HandleIndex)
    e.POST("/name", endpoints.HandleCreateName)
    e.DELETE("/name/:id", endpoints.HandleDeleteName)

    e.Logger.Fatal(e.Start(":42069"))
}

