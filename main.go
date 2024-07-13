package main

import (
    "html/template"
    "io"
    "log"
   

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "simplefin/handlers"
)

// TemplateRegistry is a custom html/template renderer for Echo framework
type TemplateRegistry struct {
    templates *template.Template
}

// Render renders a template document
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    log.Printf("Rendering template: %s", name) // Log template rendering
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Static files
    e.Static("/static", "static")

    // Templates
    t := &TemplateRegistry{
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
    e.Renderer = t

    // Routes
    e.GET("/", handlers.Index)
    e.GET("/budget", handlers.Budget)

    // Start server
    log.Println("Starting server on :8080")
    e.Logger.Fatal(e.Start(":8080"))
}
