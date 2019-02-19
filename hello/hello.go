package hello

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var debug bool
var interactive bool

// Hello says Hi
func Hello(args []string) {
	interactive, _ = strconv.ParseBool(args[0])
	debug, _ = strconv.ParseBool(args[1])
	if debug {
		fmt.Println("DEBUG: Preparing to greet the world")
	}
	fmt.Println("Hello World")
	if debug {
		fmt.Println("DEBUG: I EXIST!")
	}
	if debug {
		fmt.Printf("DEBUG: Got args: %v with len of %v\n", args, len(args))
	}
	if interactive {
		interaction(args[2])
	}
}

func interaction(server string) {
	if debug {
		fmt.Println("DEBUG: Entering interactive mode")
	}
	for {
		fmt.Printf("[GN]:")
		reader := bufio.NewReader(os.Stdin)
		inputString := readFromStdin(reader)
		switch inputString {
		case "hi":
			fmt.Println("Nice to meet you!")
		case "what is your name?":
			fmt.Println("I don't know yet...")
			time.Sleep(1 * time.Second)
			fmt.Println("hrm...")
			time.Sleep(1 * time.Second)
			fmt.Println("I am groot?")
		case "goodbye":
			fmt.Println("Bye!")
			if debug {
				fmt.Println("DEBUG: Exiting")
			}
			os.Exit(0)
		case "groot":
			fmt.Println("I am groot")
		case "connect":
			if debug {
				fmt.Printf("DEBUG: Attempting to establish connection to %s...\n", server)
			}
			fmt.Printf("Beep boop...connecting to %s...\n", server)
			time.Sleep(1 * time.Second)
			fmt.Println("Oops! That didn't work.")
			if debug {
				fmt.Printf("DEBUG: Unable to establish connection to %s: network unavailable to groot\n", server)
			}
		default:
			if debug {
				fmt.Println("DEBUG: Unable to parse user input")
			}
			fmt.Println("Sorry I don't understand you.")
			fmt.Println("You can ask me the following:")
			fmt.Println(" - 'hi'")
			fmt.Println(" - 'what is your name?'")
			fmt.Println(" - 'groot'")
			fmt.Println(" - 'connect'")
			fmt.Println(" - 'goodbye'")
			fmt.Println("\nI'm afraid that's all I know!")
		}
	}
}

func readFromStdin(reader *bufio.Reader) string {
	in, _ := reader.ReadBytes('\n')
	in = in[:len(in)-1]
	return string(in)
}
