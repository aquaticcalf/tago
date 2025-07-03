package tago

import (
	"fmt"
	"html"
	"io"
)

// Node is anything that can render itself into HTML
type Node interface {
	Render(w io.Writer) error
}

// Element represents an HTML element with tag, attributes, and children
type Element struct {
	tag      string
	attrs    map[string]string
	children []Node
}

// Option configures an Element ( attributes or children )
type Option func(*Element)

// NewElement constructs an Element for the given tag, applying opts
func NewElement(tag string, opts ...Option) *Element {
	e := &Element{
		tag:   tag,
		attrs: make(map[string]string),
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

// Attr sets an attribute key="value" on an Element
func Attr(key, value string) Option {
	return func(e *Element) {
		e.attrs[key] = value
	}
}

// Child appends one or more child Nodes to an Element
func Child(nodes ...Node) Option {
	return func(e *Element) {
		e.children = append(e.children, nodes...)
	}
}

// Render outputs the HTML for the Element to w
func (e *Element) Render(w io.Writer) error {
	// Opening tag and attributes
	if _, err := fmt.Fprintf(w, "<%s", e.tag); err != nil {
		return err
	}
	for k, v := range e.attrs {
		if _, err := fmt.Fprintf(w, ` %s="%s"`, k, html.EscapeString(v)); err != nil {
			return err
		}
	}

	// Void element : close immediately
	if isVoidElement(e.tag) {
		// HTML5 void tags should not have a closing tag
		_, err := fmt.Fprint(w, ">")
		return err
	}

	// Non void : finish opening, render children, then closing
	fmt.Fprint(w, ">")
	for _, c := range e.children {
		if err := c.Render(w); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(w, "</%s>", e.tag); err != nil {
		return err
	}

	return nil
}

// Text represents plain text content ( HTML escaped )
type Text string

// Render writes escaped text
func (t Text) Render(w io.Writer) error {
	_, err := w.Write([]byte(html.EscapeString(string(t))))
	return err
}

// Raw represents unescaped HTML content
type Raw string

// Render writes raw HTML
func (r Raw) Render(w io.Writer) error {
	_, err := w.Write([]byte(r))
	return err
}
