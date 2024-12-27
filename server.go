package main

import (
	"go-server/handlers"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	e := echo.New()
	renderer := &Template{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.GET("/users/search", handlers.SearchUser)
	e.GET("/add-user", func(c echo.Context) error {
		return c.Render(http.StatusOK, "add-user", map[string]interface{}{})
	})
	e.POST("/users", handlers.AddUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)
	
	e.Logger.Fatal(e.Start(":1400"))
}