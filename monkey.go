package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/negativ/monkey/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Starting Monkey REPL...\n", user.Name)

	repl.Start(os.Stdin, os.Stdout)
}
