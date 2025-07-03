package tago

// Class sets the class attribute on an Element
func Class(name string) Option {
	return Attr("class", name)
}

// ID sets the ID attribute on an Element
func ID(name string) Option {
	return Attr("id", name)
}
