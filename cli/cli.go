package cli

import (
	"strings"
	"fmt"
	"os"

	"github.com/grellyd/goNew/hello"
)


// Run the cli
func Run() (code int) {
	commands := ValidCommands()
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for _, cmd := range commands {
			arg := getFullArg(args, i, cmd.length)
			if cmd.regex.MatchString(arg) {
				if cmd.short == "-h" {
					fmt.Println(usage(commands))
					return code
				}
				cmd.present = true
			}
		}
	}
	///----
	fmt.Println("Commands Present")
	for _, cmd := range commands {
		fmt.Printf("%s: %v\n", cmd.long, cmd.present)
	}
	///----
	err := runPackage(commands)
	if err != nil {
		code = 1
	}
	return code
}

func getFullArg(args []string, start int, length int) string {
	if length == 1 || len(args) < (start+length) {
		return args[start]
	}
	return strings.Join(args[start:start + length], " ")
}

// call the main package with the given commands
func runPackage(commands []*Command) error {
	args := convertCommands(commands)
	hello.Hello(args)
	return nil
}

func usage(commands []*Command) (usageString string){
	var builder strings.Builder
	builder.WriteString(breakLine)
	builder.WriteString("\n")
	builder.WriteString(programName)
	builder.WriteString("\n")
	builder.WriteString(breakLine)
	builder.WriteString("\n")
	builder.WriteString(programDescription)
	builder.WriteString("\n\n")
	fmt.Fprintf(&builder, "Usage: %s [options]\n\n", usageCommand)
	fmt.Fprint(&builder, "Valid options:\n\n")
	for _, cmd := range(commands) {
		fmt.Fprintf(&builder, "%s/%s: %s\n", cmd.short, cmd.long, cmd.description)
	}
	return builder.String()
}

func convertCommands(commands []*Command) (args []string) {
	args = make([]string, len(commands))
	for _, cmd := range commands {
		if cmd.present {
			args[cmd.order] = "true"
		}
	}
	return args
}
