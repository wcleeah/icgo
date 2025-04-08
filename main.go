package main 

import (
	"fmt"
	"os"
	"os/user"

	"com.lwc.icgo/repl"
)

func main() {
	u, err := user.Current()

	if err != nil {
		panic("random error when getting os user")
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
