package main

import (
	"flag"
	"go-echo/internal/web"
)

const dirPath = "/Users/lijia/go/src/go-echo/data/"
const nodePath = ""

func main() {
	dirPtr := flag.String("d", dirPath, "Path to config dir")
	nodePtr := flag.String("n", nodePath, "Path to node modules")
	flag.Parse()

	web.Gui(*dirPtr, *nodePtr)
}
