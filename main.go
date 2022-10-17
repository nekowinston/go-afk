package main

import (
	"embed"
	"flag"
	"io/fs"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

//go:embed web/dist
var embeddedFiles embed.FS

func getFileSystem() http.FileSystem {
	fsys, _ := fs.Sub(embeddedFiles, "web/dist")
	return http.FS(fsys)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		rateLimit = 20
	}
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(rateLimit))))

	e.GET("/", imageGenerator)

	assetHandler := http.FileServer(getFileSystem())
	e.GET("/*", echo.WrapHandler(assetHandler))

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
	params := c.QueryParams()

	if len(params.Encode()) == 0 {
		index, _ := embeddedFiles.ReadFile("web/dist/index.html")
		return c.HTMLBlob(200, index)
	}

	var options Options
	if params.Has("t") {
		options.Text = c.QueryParam("t")
	} else {
		options.Text = "Hello, World!"
	}
	if params.Has("f") {
		options.Flavor = c.QueryParam("f")
	} else {
		options.Flavor = "frappe"
	}
	if params.Has("c") {
		options.Color = c.QueryParam("c")
	} else {
		options.Color = "pink"
	}
	options.Italic = params.Has("i")

	img, err := TextOnImg(options)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(http.StatusOK, "image/png", img.Bytes())
}
