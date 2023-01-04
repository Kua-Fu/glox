package main

import (
	"fmt"
	"learning/glox/scanner"
	"os"
)

func main() {
	args := os.Args
	// 1. 如果传参大于1个，报错
	if len(args) > 2 {
		fmt.Println("usage: glox [script]")
		return
	}
	scanner.StartScanner(args)
}
