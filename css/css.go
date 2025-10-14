package css

import "aqclf.xyz/tago"

func Daisy() *tago.Element {
	return tago.Fragment(
		tago.Link(
			tago.Attr("href", "https://cdn.jsdelivr.net/npm/daisyui@5"),
			tago.Attr("rel", "stylesheet"),
		),
		tago.Script(tago.Attr("src", "https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4")),
		tago.Link(
			tago.Attr("href", "https://cdn.jsdelivr.net/npm/daisyui@5/themes.css"),
			tago.Attr("rel", "stylesheet"),
		),
	)
}

func Tailwind() *tago.Element {
	return tago.Script(
		tago.Attr("src", "https://cdn.tailwindcss.com"),
	)
}
