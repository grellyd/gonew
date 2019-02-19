package cli

import (
	"regexp"
)

// A Command is an application instruction. 
// Each must either correspond to a parameter in the main program, or act in the cli
type Command struct {
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
		long: "--interactive",
		short: "-i",
		description: "Run the program in interactive mode",
		regex: *regexp.MustCompile("( -i )|( --interactive)"),
	}
}

func debug() *Command {
	return &Command{
		long: "--debug",
		short: "-d",
		description: "Run the program in debug mode, with extra logging",
		regex: *regexp.MustCompile("( -d )|( --debug )"),
	}
}

func help() *Command {
	return &Command{
		long: "--help",
		short: "-h",
		description: "Display this help info",
		regex: *regexp.MustCompile("( -h )|( --help )"),
	}
}
