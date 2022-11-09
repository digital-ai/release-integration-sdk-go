package task

import "encoding/json"

type PropertyDefinition struct {
	Name       string          `json:"name"`
	Value      json.RawMessage `json:"value"`
	Kind       string          `json:"kind"`
	Category   string          `json:"category"`
	IsPassword bool            `json:"isPassword"`
}

type TaskContext struct {
	Id         string               `json:"id"`
	Type       string               `json:"type"`
	Properties []PropertyDefinition `json:"properties"`
}

type AutomatedTaskAsUserContext struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReleaseContext struct {
	Id                  string                     `json:"id"`
	AutomatedTaskAsUser AutomatedTaskAsUserContext `json:"automatedTaskAsUser"`
}

type InputContext struct {
	Release ReleaseContext `json:"release"`
	Task    TaskContext    `json:"task"`
}

type TaskResult struct {
	Response                string            `json:"response"`
	ResponseInt             int64             `json:"responseInt"`
	ResponseBool            bool              `json:"responseBool"`
	ResponseDate            string            `json:"responseDate"`
	ResponseEnum            string            `json:"responseEnum"`
	ResponseListString      []string          `json:"responseListString"`
	ResponseSetString       []string          `json:"responseSetString"`
	ResponseMapStringString map[string]string `json:"responseMapStringString"`
	UnknownProperty         string            `json:"unknownProperty"`
}

type TaskOutputContext struct {
	ExitCode         int64      `json:"exitCode"`
	OutputProperties TaskResult `json:"outputProperties"`
}
