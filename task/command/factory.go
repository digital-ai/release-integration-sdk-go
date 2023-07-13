package command

// CommandFactory represents a factory for creating commands.
type CommandFactory interface {
	// InitCommand initializes and returns a CommandExecutor for the specified command type.
	// It returns an error if the command type is not supported.
	InitCommand(commandType CommandType) (CommandExecutor, error)
}
