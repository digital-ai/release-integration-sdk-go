package util

import (
	"encoding/json"
	"fmt"
	"github.com/savaki/jq"
	"time"
)

// ParseDate parses a date string into the specified sample format.
func ParseDate(sampleFormat string, dateTime string) (string, error) {
	parsedDateTime, err := time.Parse(sampleFormat, dateTime)
	if err != nil {
		return "", fmt.Errorf("error parsing date: %v", err)

	}
	return parsedDateTime.Format(time.RFC3339), nil
}

// ParseNode applies a JQ operation on the JSON payload.
func ParseNode(jqOp string, result json.RawMessage) ([]byte, error) {
	parse, err := jq.Parse(jqOp)
	if err != nil {
		return nil, fmt.Errorf("could not create parser for JQ operation '%s': %v", jqOp, err)
	}
	parseResult, err := parse.Apply(result)
	if err != nil {
		return nil, fmt.Errorf("could not apply parser for JQ operation '%s': %v", jqOp, err)
	}
	return parseResult, nil
}
