package main

import (
	"fmt"
	"learning/glox/parser"
	"os"
)

func main() {

	args := os.Args
	// 1. if args large than one, print error
	if len(args) > 2 {
		fmt.Println("usage: glox [script]")
		return
	}
	parser.StartParse(args)

}
