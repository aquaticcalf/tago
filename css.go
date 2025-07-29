package tago

func Daisy() *Element {
	return Fragment(
		Link(
			Attr("href", "https://cdn.jsdelivr.net/npm/daisyui@5"),
			Attr("rel", "stylesheet"),
		),
		Script(Attr("src", "https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4")),
		Link(
			Attr("href", "https://cdn.jsdelivr.net/npm/daisyui@5/themes.css"),
			Attr("rel", "stylesheet"),
		),
	)
}

func Tailwind() *Element {
	return Script(
		Attr("src", "https://cdn.tailwindcss.com"),
	)
}
