package handlers

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
    Templates *template.Template
}

func NewTemplate(pattern string) (*Template, error) {
    tmpl, err := template.ParseGlob(pattern)
    if err != nil {
        return nil, err
    }
    return &Template{
        Templates: tmpl,
    }, nil
}

// Render рендерит шаблон
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.Templates.ExecuteTemplate(w, name, data)
}