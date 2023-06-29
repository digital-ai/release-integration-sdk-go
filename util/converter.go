package util

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

// YamlToJson converts YAML data to JSON.
func YamlToJson(yamlData []byte) ([]byte, error) {
	var data interface{}
	err := yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
