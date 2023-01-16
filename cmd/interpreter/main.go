package main

import (
	"fmt"
	"learning/glox/interpreter"
	"os"
)

func main() {
	args := os.Args
	// 1. if args large than one, print error
	if len(args) > 2 {
		fmt.Println("usage: glox [script]")
		return
	}
	interpreter.StartInterpreter(args)
}
