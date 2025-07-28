package main

import (
	. "aqclf.xyz/tago"
)

func main() {
	page := Html(
		Head(
			Tailwind(),
			Title("tago"),
			Link(
				Attr("rel", "shortcut icon"),
				Attr("href", "https://aqclf.xyz/favicon.png"),
				Attr("type", "image/png"),
			),
			Meta(
				Attr("name", "viewport"),
				Attr("content", "width=device-width, initial-scale=1.0"),
			),
		),
		Body(
			Class("bg-black text-white min-h-[100dvh] flex flex-col items-center justify-center"),
			H1(Class("text-4xl"), "Hello, world!"),
		),
	)

	Build(page, "index")
}
