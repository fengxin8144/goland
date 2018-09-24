package main

import (
	"flag"
	"fmt"
)

func main() {
	var src string
	var level *int
	var memo string
	flag.Parse()
	flag.Usage()
	fmt.Printf("src=%s, level=%d, memo=%s\n", src, *level, memo)
}
