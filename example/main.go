package main

import "github.com/aquaticcalf/tago"

func main() {
	page := tago.Html(
		tago.Child(tago.Head(
			tago.Child(
				tago.Script(
					tago.Attr("src", "https://cdn.tailwindcss.com"),
				),
			),
			tago.Child(tago.Title(
				tago.Child(tago.Text("tago")),
			)),
			tago.Child(tago.Link(
				tago.Attr("rel", "shortcut icon"),
				tago.Attr("href", "https://aqclf.xyz/favicon.png"),
				tago.Attr("type", "image/png"),
			)),
			tago.Child(tago.Meta(
				tago.Attr("name", "viewport"),
				tago.Attr("content", "width=device-width, initial-scale=1.0"),
			)),
		)),
		tago.Child(
			tago.Body(
				tago.Class("min-h-[100dvh] flex flex-col items-center justify-center bg-gray-900 text-white"),
				tago.Child(
					tago.H1(
						tago.Class("text-4xl font-semibold"),
						tago.Child(tago.Text("tago")),
					),
				),
			),
		),
	)

	app := tago.NewApp()
	app.GET("/", page)
	app.Run(":8080")
}
