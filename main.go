package main

import (
	"bytes"
	"flag"
	"net/http"

	"github.com/chai2010/webp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", imageGenerator)

	port := flag.String("port", "3000", "port to listen on")
	e.Logger.Fatal(e.Start(":" + *port))
}

type Options struct {
	Text   string
	Flavor string
	Color  string
	Italic bool
}

func imageGenerator(c echo.Context) error {
	text := c.QueryParam("t")
	flavor := c.QueryParam("f")
	color := c.QueryParam("c")
	italic := c.QueryParam("i")
	if text == "" {
		text = "Hello, World!"
	}
	// if flavour isn't of type flavour
	if flavor == "" {
		flavor = "frappe"
	}
	if color == "" {
		color = "pink"
	}

	img, err := TextOnImg(
		Options{
			Text:   text,
			Flavor: flavor,
			Color:  color,
			Italic: italic != "",
		},
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	buffer := new(bytes.Buffer)
	if err := webp.Encode(buffer, img, &webp.Options{Quality: 95}); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(http.StatusOK, "image/webp", buffer.Bytes())
}
