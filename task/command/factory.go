package command

type CommandFactory interface {
	InitCommand(commandType CommandType) (CommandExecutor, error)
}
