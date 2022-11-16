package command

import (
	"encoding/json"
)

type CommandExecutor interface {
	Execute() (json.RawMessage, bool, error)
}
