package cli

import (
	"strings"
	"fmt"
	"os"

	"github.com/grellyd/goNew/hello"
)

const (
	breakLine = "====================\n"
	programName = "goNew"
	usageCommand = "go run goNew.go"

)


// Run the cli
func Run() (code int) {
	commands := ValidCommands()
	args := os.Args[1:]
	for _, arg := range args {
		for _, cmd := range commands {
			if cmd.regex.MatchString(arg) {
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

// call the main package with the given commands
func runPackage(commands []*Command) error {
	args := convertCommands(commands)
	hello.Hello(args)
	return nil
}

func usage(commands []*Command) (usageString string){
	var builder strings.Builder
	builder.WriteString(breakLine)
	builder.WriteString(programName)
	builder.WriteString(breakLine)
	fmt.Fprintf(&builder, "Usage: %s [options]\n", usageCommand)
	fmt.Fprint(&builder, "Valid options:\n\n")
	for _, cmd := range(commands) {
		fmt.Fprintf(&builder, "%s/%s: %s", cmd.long, cmd.short, cmd.description)
	}
	return builder.String()
}

func convertCommands(commands []*Command) []string {
	return nil
}
