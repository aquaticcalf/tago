package js

func log(thing string) string {
    return `console.log("` + thing + `")`
}

func alert(thing string) string {
	return `alert("` + thing + `")`
}

// no, this is not how i want it to be :/
// it should be auto generated, i am thinking of using antlr
// i need a bridge from go to javascript