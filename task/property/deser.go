package property

import (
	"encoding/json"
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"k8s.io/klog"
)

const (
	Null = "null"
)

func ExtractType(fieldLabel string, properties []task.PropertyDefinition, extract any) error {
	server, err := ExtractByName(fieldLabel, properties)
	if err != nil {
		klog.Errorf("Cannot extract field %s: %v", fieldLabel, err)
		return err
	}
	return DeserializeType(fieldLabel, server, extract)
}

func DeserializeType(fieldLabel string, rawProperty json.RawMessage, extract any) error {
	serverProperties, err := Deserialize(rawProperty)
	if err != nil {
		klog.Errorf("Cannot deserialize properties for field %s: %v", fieldLabel, err)
		return err
	}

	serverProps := make(map[string]json.RawMessage)
	for _, data := range serverProperties {
		serverProps[data.Name] = data.Value
	}
	serverPropsJson, err := json.Marshal(serverProps)
	if err != nil {
		klog.Errorf("Cannot process properties entry %v", err)
		return err
	}
	unMarshalErr := json.Unmarshal(serverPropsJson, extract)
	if unMarshalErr != nil {
		klog.Errorf("Cannot unmarshal field to struct %v", unMarshalErr)
		return err
	}
	return nil
}

func ExtractByName(field string, properties []task.PropertyDefinition) (json.RawMessage, error) {
	for _, property := range properties {
		if property.Name == field {
			if !json.Valid(property.Value) {
				return nil, fmt.Errorf("cannot deserialize invalid json")
			}
			if string(property.Value) == Null {
				return nil, fmt.Errorf("cannot deserialize json null value")
			}
			return property.Value, nil
		}
	}
	return nil, fmt.Errorf("cannot find field %s in properties", field)
}

func Deserialize(raw json.RawMessage) ([]task.PropertyDefinition, error) {
	var propertyWrapper task.ComposedProperty
	unMarshalErr := json.Unmarshal(raw, &propertyWrapper)
	if unMarshalErr != nil {
		klog.Errorf("Cannot umarshal properties: %v", unMarshalErr)
		return nil, fmt.Errorf("failed to unmarshal properties: %w", unMarshalErr)
	}
	return propertyWrapper.Properties, nil
}