package tago

// voidElements holds HTML5 void ( self closing ) tags
var voidElements = map[string]struct{}{
	"area": {}, "base": {}, "br": {}, "col": {}, "command": {},
	"embed": {}, "hr": {}, "img": {}, "input": {}, "keygen": {},
	"link": {}, "meta": {}, "param": {}, "source": {}, "track": {},
	"wbr": {},
}

// isVoidElement reports whether a tag is a void element
func isVoidElement(tag string) bool {
	_, ok := voidElements[tag]
	return ok
}
