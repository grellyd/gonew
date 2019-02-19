package hello

import "fmt"

// Hello says Hi
func Hello(args []string) {
	fmt.Println("Hello World")
	fmt.Printf("Got args: %v\nWith len of %v\n", args, len(args))
}

// TODO: Add interactive and debug functionality
