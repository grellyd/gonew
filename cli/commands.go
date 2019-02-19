package cli

import (
	"regexp"
)

const (
	breakLine = "===================="
	programName = "goNew"
	usageCommand = "go run goNew.go"

)

// A Command is an application instruction. 
type Command struct {
	order int
	long string
	short string
	description string
	regex regexp.Regexp
	present bool
}

// ValidCommands to this program
func ValidCommands() (commands []*Command) {
	return []*Command {debug(), interactive(), help()}
}

func interactive() *Command {
	return &Command{
		order: 0,
		long: "--interactive",
		short: "-i",
		description: "Run the program in interactive mode",
		regex: *regexp.MustCompile("-i|--interactive"),
	}
}

func debug() *Command {
	return &Command{
		order: 1,
		long: "--debug",
		short: "-d",
		description: "Run the program in debug mode, with extra logging",
		regex: *regexp.MustCompile("-d|--debug"),
	}
}

func help() *Command {
	return &Command{
		// Help is handled by the cli wrapper
		order: -1,
		long: "--help",
		short: "-h",
		description: "Display this help dialogue and exit",
		regex: *regexp.MustCompile("-h|--help"),
	}
}
