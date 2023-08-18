package main

import (
	"html/template"
	"log"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"theprimeagen.tv/filter_simple_list/pkg/database"
	"theprimeagen.tv/filter_simple_list/pkg/endpoints"
)

func main() {
    err := database.InitContacts("file:///tmp/contacts")
    if err != nil {
        log.Fatalf("could not initialize db: %+v", err)
    }

    tmpl, err := template.ParseFiles(
        "./public/views/contacts/index.html",
        "./public/views/contacts/help.html",
        "./public/views/contacts/settings.html",
        "./public/views/contacts/header.html",
        "./public/views/contacts/contacts.html",
    )
    if err != nil {
        log.Fatalf("could not initialize templates: %+v", err)
    }

    e := echo.New()
    e.Renderer = endpoints.NewTemplateRenderer(tmpl)
    e.Use(middleware.Logger())
    e.Static("/css", "css");

    e.GET("/", endpoints.HandleIndex)

    e.GET("/contacts", endpoints.HandleIndex)

    e.GET("/contacts/new", endpoints.HandleNewContact)
    e.POST("/contacts/new", endpoints.HandleCreateContact)
    e.DELETE("/contacts/:id", endpoints.HandleDeleteContact)
    e.GET("/contacts/:id", endpoints.HandleEditContact)
    e.PUT("/contacts/:id", endpoints.HandleCreateContact)
    e.POST("/contacts/email", endpoints.HandleValidateEmail)

    e.GET("/settings", endpoints.HandleSettings)
    e.GET("/help", endpoints.HandleHelp)

    e.Logger.Fatal(e.Start(":42069"))
}

