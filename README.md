# goNew
## An Intelligent Go Program CLI wrapper
### By @grellyd

goNew is a base cli interface for an arbitrary go program. The base provides automatic argument checking, help usage, and an importable structure for the core code. 

In this package, 'hello' is the arbitrary go program, which just prints "hello world". It also accepts a debug and interactive flag, which respectively change the behaviour of the program. As an example for more complex flags, it accepts a server argument with an IP:Port pair.

When building a new package, in order to get these features, you must register your commands with the cli package. As you add command line arguments, add a `*Command` in `cli/commands.go` such that `ValidCommands()` returns the complete set of commands for the program. This is the only file you have to modify. Further for a detailed help flag, fill in the constants at the top of the `cli/commands.go` file.

Note that the order given for each command my correspond to the order in which your core program expects the arguments. While this does break the law of demeter, as the cli package has some knowledge of how the core package functions, the cli package cannot act on arbitrary input without some knowledge of what to expect.
