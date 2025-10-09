package tago

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// EchoHandler creates an Echo route handler that renders the given Node
func EchoHandler(node Node) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Set HTML content type
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		// Render the HTML
		if err := node.Render(c.Response().Writer, 0); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

// App wraps an Echo server with Tago routes
type App struct{ e *echo.Echo }

// NewApp creates a new Tago App
func NewApp() *App { return &App{e: echo.New()} }

// GET registers a GET route serving a Node
func (a *App) GET(path string, n Node) {
	a.e.GET(path, EchoHandler(n))
}

// Run starts the HTTP server
func (a *App) Run(addr string) error {
	return a.e.Start(addr)
}
