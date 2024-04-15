package ci

import (
	"encoding/json"
	"reflect"
	"time"
)

type TypeMap map[reflect.Type]string

var typeMapRegistry = make(TypeMap)

func RegisterTypeMappings(types TypeMap) {
	for reflection, typeName := range types {
		typeMapRegistry[reflection] = typeName
	}
}

func RegisterTypeMapping(typeReflection reflect.Type, typeName string) {
	typeMapRegistry[typeReflection] = typeName
}

func ToSyntheticType(searchedType reflect.Type) string {
	if typeName, ok := typeMapRegistry[searchedType]; ok {
		return typeName
	}
	return searchedType.Name()
}

// SerializeStruct serializes any struct into the desired JSON format.
func SerializeCi(data interface{}) ([]byte, error) {
	jsonData, err := serializeData(data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonData)
}

func serializeData(data interface{}) (map[string]interface{}, error) {
	jsonData := make(map[string]interface{})
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	jsonData["id"] = nil
	jsonData["type"] = ToSyntheticType(dataType)
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		tag := field.Tag.Get("synthetic")
		fieldValue := dataValue.Field(i)

		if len(tag) > 0 {
			// Handle synthetic fields
			switch fieldValue.Kind() {
			case reflect.Interface:
				nestedJSON, err := serializeData(fieldValue.Interface())
				if err != nil {
					return nil, err
				}
				jsonData[tag] = nestedJSON
			case reflect.Struct:
				if field.Type == reflect.TypeOf(time.Time{}) {
					// Convert time.Time to ISO format string
					jsonData[tag] = fieldValue.Interface().(time.Time).Format(time.RFC3339)
				} else {
					nestedJSON, err := serializeData(fieldValue.Interface())
					if err != nil {
						return nil, err
					}
					jsonData[tag] = nestedJSON
				}
			default:
				if !isEmptyValue(fieldValue) {
					value := fieldValue.Interface()
					jsonData[tag] = value
				}
			}
		} else if field.Anonymous {
			// Handle embedded fields
			embedded, err := serializeData(fieldValue.Interface())
			if err != nil {
				return nil, err
			}
			for k, v := range embedded {
				jsonData[k] = v
			}
		}
	}

	return jsonData, nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(time.Time{}) {
			return v.Interface().(time.Time).IsZero()
		}
	default:
		return v.IsNil()
	}
	return false
}
