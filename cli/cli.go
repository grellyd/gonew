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
	err := runPackage(commands)
	if err != nil {
		code = 1
	}
	return code
}

// call the main package with the given commands
func runPackage(commands []*Command) error {
	
	config := hello.Config(commands) 
	hello.Hello(configuration)
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
		fmt.Fprintf(&builder, "%s/%s: %s", cmd.longFlag, cmd.shortFlag, cmd.description)
	}
	return builder.String()
}


