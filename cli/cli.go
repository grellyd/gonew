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
	for _, arg := range args {
		for _, cmd := range commands {
			if cmd.regex.MatchString(arg) {
				cmd.present = true
				if cmd.short == "-h" {
					fmt.Println(usage(commands))
					return code
				}
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
	fmt.Fprintf(&builder, "Usage: %s [options]\n", usageCommand)
	fmt.Fprint(&builder, "Valid options:\n\n")
	for _, cmd := range(commands) {
		fmt.Fprintf(&builder, "%s/%s: %s\n", cmd.long, cmd.short, cmd.description)
	}
	return builder.String()
}

func convertCommands(commands []*Command) (args []string) {
	return args
}
