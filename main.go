package main

import (
	"fmt"
	"github.com/rjmccabe3701/go_examples/morestrings"
	"github.com/rjmccabe3701/go_examples/my_util"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println("filename = " + argsWithoutProg[0])
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	my_util.PrintLines(argsWithoutProg[0])
}
