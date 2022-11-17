package command

import (
	"encoding/json"
)

type CommandExecutor interface {
	FetchResult() (json.RawMessage, bool, error)
}
