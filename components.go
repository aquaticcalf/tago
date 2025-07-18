package tago

func Html(items ...any) *Element     { return NewElement("html", items...) }
func Head(items ...any) *Element     { return NewElement("head", items...) }
func Title(items ...any) *Element    { return NewElement("title", items...) }
func Meta(items ...any) *Element     { return NewElement("meta", items...) }
func Link(items ...any) *Element     { return NewElement("link", items...) }
func Body(items ...any) *Element     { return NewElement("body", items...) }
func Script(items ...any) *Element   { return NewElement("script", items...) }
func Style(items ...any) *Element    { return NewElement("style", items...) }
func Div(items ...any) *Element      { return NewElement("div", items...) }
func Span(items ...any) *Element     { return NewElement("span", items...) }
func P(items ...any) *Element        { return NewElement("p", items...) }
func H1(items ...any) *Element       { return NewElement("h1", items...) }
func H2(items ...any) *Element       { return NewElement("h2", items...) }
func H3(items ...any) *Element       { return NewElement("h3", items...) }
func Ul(items ...any) *Element       { return NewElement("ul", items...) }
func Ol(items ...any) *Element       { return NewElement("ol", items...) }
func Li(items ...any) *Element       { return NewElement("li", items...) }
func A(items ...any) *Element        { return NewElement("a", items...) }
func Img(items ...any) *Element      { return NewElement("img", items...) }
func Button(items ...any) *Element   { return NewElement("button", items...) }
func Form(items ...any) *Element     { return NewElement("form", items...) }
func Input(items ...any) *Element    { return NewElement("input", items...) }
func Label(items ...any) *Element    { return NewElement("label", items...) }
func Textarea(items ...any) *Element { return NewElement("textarea", items...) }
func Select(items ...any) *Element   { return NewElement("select", items...) }
func Table(items ...any) *Element    { return NewElement("table", items...) }
func Thead(items ...any) *Element    { return NewElement("thead", items...) }
func Tbody(items ...any) *Element    { return NewElement("tbody", items...) }
func Tr(items ...any) *Element       { return NewElement("tr", items...) }
func Td(items ...any) *Element       { return NewElement("td", items...) }
func Th(items ...any) *Element       { return NewElement("th", items...) }
