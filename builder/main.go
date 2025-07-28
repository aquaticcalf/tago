package main

import (
	"flag"
	"fmt"
)

// i need to read the file system in pages folder recursively and then generate corresponding folder/index.html structure..
// it shall work like https://npmjs.com/package/file-system-router

func main() {
	fmt.Println("welcome to tago builder")

	pagesFolderFlag := flag.String("folder", "pages", "Pages Folder")

	flag.Parse()

	pagesFolder := *pagesFolderFlag

	fmt.Println(pagesFolder)
}

// i do not know how to make this work yet, i need to figure out how i want to get the "default export" from each page
// my present "best" idea is to get it via AST