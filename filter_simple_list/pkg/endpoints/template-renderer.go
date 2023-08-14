package endpoints

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
    tmpl *template.Template
}

func NewTemplateRenderer(tmpls *template.Template) TemplateRenderer {
    return TemplateRenderer{tmpls}
}

func (t TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.tmpl.ExecuteTemplate(w, name, data)
}
