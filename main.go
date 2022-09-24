package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func mainHandler(c echo.Context) error {
	adres := c.Request().RemoteAddr
	real := c.RealIP()
	fmt.Println(real)
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "merd!",
		"ip":   adres,
	})
}

var sms string = "sms"

func ipHandler(c echo.Context) error {
	ip := c.QueryParam("ip")
	println(ip)
	/*	host := c.Request().Host
		adres := c.Request().RemoteAddr*/
	if a == true {
		println("asddsa")

		a = false

		return c.String(http.StatusOK, sms)
	}
	return c.String(http.StatusOK, ip)

}

var a bool

func redirectPage(c echo.Context) error {
	adres := c.Request().RemoteAddr
	kod := c.QueryParam("kod")
	fmt.Println(kod)
	if kod == "" {
		fmt.Println("test")
	} else {
		fmt.Println("kodun ici dolu:", kod)
		a = true

	}
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "merd!",
		"ip":   adres,
		"kod":  kod,
	})
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/", mainHandler)
	e.POST("/islem", ipHandler)
	e.GET("/islem", ipHandler)
	e.GET("/iste", redirectPage)

	e.Start(":70")

}
