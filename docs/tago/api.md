# tago api

## types

### node
interface for anything that can render itself into html

```go
type Node interface {
    Render(w io.Writer, indent int) error
}
```

### element
represents an html element with tag, attributes, and children

```go
type Element struct {
    // fields are unexported
}
```

### text
represents plain text content html escaped

```go
type Text string
```

### raw
represents unescaped html content

```go
type Raw string
```

### option
configures an element attributes or children

```go
type Option interface {
    apply(*Element)
}
```

## functions

### newelement
constructs an element for the given tag, applying any mix of options or nodes

```go
el := NewElement("div", Attr("id", "mydiv"), Text("hello"))
```

### attr
sets an attribute key="value" on an element

```go
el := NewElement("a", Attr("href", "https://example.com"))
```

### child
appends one or more child nodes to an element

```go
el := NewElement("ul", Child(
    NewElement("li", Text("item 1")),
    NewElement("li", Text("item 2")),
))
```

### class
sets the class attribute on an element

```go
el := NewElement("div", Class("container"))
```

### id
sets the id attribute on an element

```go
el := NewElement("div", ID("main"))
```

### fragment
creates a fragment element with no tag

```go
frag := Fragment(
    NewElement("h1", Text("title")),
    NewElement("p", Text("content")),
)
```

### build
builds the page to a file

```go
page := Html(Head(Title("page")), Body(H1("hello")))
Build(page, "index")
```

## components

### html
creates html element

```go
page := Html(Head(Title("title")), Body(P("content")))
```

### head
creates head element

```go
head := Head(Title("title"), Meta(Attr("name", "description")))
```

### title
creates title element

```go
title := Title("my page")
```

### meta
creates meta element

```go
meta := Meta(Attr("name", "viewport"), Attr("content", "width=device-width"))
```

### link
creates link element

```go
link := Link(Attr("rel", "stylesheet"), Attr("href", "style.css"))
```

### body
creates body element

```go
body := Body(Div(Class("content"), Text("hello")))
```

### script
creates script element

```go
script := Script(Attr("src", "script.js"))
```

### style
creates style element

```go
style := Style(Raw("body { color: red; }"))
```

### div
creates div element

```go
div := Div(Class("wrapper"), Text("content"))
```

### span
creates span element

```go
span := Span(Class("highlight"), Text("text"))
```

### p
creates p element

```go
p := P("paragraph text")
```

### h1
creates h1 element

```go
h1 := H1("heading")
```

### h2
creates h2 element

```go
h2 := H2("subheading")
```

### h3
creates h3 element

```go
h3 := H3("subsubheading")
```

### ul
creates ul element

```go
ul := Ul(Li("item 1"), Li("item 2"))
```

### ol
creates ol element

```go
ol := Ol(Li("first"), Li("second"))
```

### li
creates li element

```go
li := Li("list item")
```

### a
creates a element

```go
a := A(Attr("href", "https://example.com"), Text("link"))
```

### img
creates img element

```go
img := Img(Attr("src", "image.jpg"), Attr("alt", "description"))
```

### button
creates button element

```go
button := Button(Attr("type", "submit"), Text("click me"))
```

### form
creates form element

```go
form := Form(Attr("action", "/submit"), Input(Attr("type", "text")))
```

### input
creates input element

```go
input := Input(Attr("type", "text"), Attr("name", "username"))
```

### label
creates label element

```go
label := Label(Attr("for", "username"), Text("username"))
```

### textarea
creates textarea element

```go
textarea := Textarea(Attr("name", "message"), Attr("rows", "4"))
```

### select
creates select element

```go
select := Select(Attr("name", "choice"), NewElement("option", Attr("value", "1"), Text("option 1")))
```

### table
creates table element

```go
table := Table(Thead(Tr(Th("header"))), Tbody(Tr(Td("data"))))
```

### thead
creates thead element

```go
thead := Thead(Tr(Th("col1"), Th("col2")))
```

### tbody
creates tbody element

```go
tbody := Tbody(Tr(Td("row1"), Td("row2")))
```

### tr
creates tr element

```go
tr := Tr(Td("cell1"), Td("cell2"))
```

### td
creates td element

```go
td := Td("table data")
```

### th
creates th element

```go
th := Th("table header")
```