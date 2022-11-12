package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/justingodden/empire/pkg/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Empire progamming Language!\n", u.Username)
	fmt.Print("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
