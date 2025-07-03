package tago

func Fragment(children ...Node) *Element {
	return &Element{
		tag:      "",
		attrs:    nil,
		children: children,
	}
}
