# echo api

## types

### app
wraps an echo server with tago routes

```go
type App struct {
    // wraps echo.Echo
}
```

## functions

### newapp
creates a new tago app

```go
import "aqclf.xyz/tago/echo"

app := echo.NewApp()
```

### get
registers a get route serving a node

```go
page := Html(Body(H1("hello")))
app.GET("/", page)
```

### run
starts the http server

```go
app.Run(":8080")
```

### echohandler
creates an echo route handler that renders the given node

```go
handler := echo.EchoHandler(page)
e.GET("/", handler)
```