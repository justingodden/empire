package main

import (
	"fmt"
	"os"

	"github.com/justingodden/empire/pkg/repl"
	"github.com/justingodden/empire/pkg/runner"
)

func main() {
	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
	} else if len(os.Args) == 2 {
		runner.Run(os.Args[1])
	} else {
		fmt.Println("[Error] Empire: Too many arguments.")
	}
}
