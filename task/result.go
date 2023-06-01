package task

import (
	"encoding/json"
	"fmt"
	"github.com/savaki/jq"
	"strconv"
	"time"
)

const (
	errorMessage = "errorMessage"
)

type Result struct {
	resultGenerators []Generator
}

func NewResult() *Result {
	return &Result{
		resultGenerators: []Generator{},
	}
}

type Generator interface {
	GenerateValue() (interface{}, error)
	FieldName() string
}

func (r *Result) addGenerator(generator Generator) *Result {
	r.resultGenerators = append(r.resultGenerators, generator)
	return r
}

// Util functions
func parseDate(sampleFormat string, dateTime string) (string, error) {
	parsedDateTime, err := time.Parse(sampleFormat, dateTime)
	if err != nil {
		return "", fmt.Errorf("error parsing date: %v", err)

	}
	return parsedDateTime.Format(time.RFC3339), nil
}

func parseNode(jqOp string, result json.RawMessage) ([]byte, error) {
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

// The ErrorGenerator - used to represent error in standardizes response
type ErrorGenerator struct {
	err error
}

func (gen ErrorGenerator) GenerateValue() (interface{}, error) {
	return gen.err.Error(), gen.err
}

func (gen ErrorGenerator) FieldName() string {
	return errorMessage
}

func (r *Result) Error(err error) *Result {
	return r.addGenerator(ErrorGenerator{err})
}

// The StringGenerator - used to represent simple string value on custom FieldName()
type StringGenerator struct {
	fieldName  string
	fieldValue string
}

func (gen StringGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

func (gen StringGenerator) FieldName() string {
	return gen.fieldName
}

func (r *Result) String(resultField string, result string) *Result {
	return r.addGenerator(StringGenerator{resultField, result})
}

// The DateGenerator - used to represent simple Date and Time value on custom FieldName()
type DateGenerator struct {
	fieldName    string
	sampleFormat string
	dateTime     string
}

func (gen DateGenerator) GenerateValue() (interface{}, error) {
	date, err := parseDate(gen.sampleFormat, gen.dateTime)
	if err != nil {
		return nil, err
	}
	return date, nil
}

func (gen DateGenerator) FieldName() string {
	return gen.fieldName
}

func (r *Result) Date(resultField string, dateTime time.Time) *Result {
	return r.addGenerator(DateGenerator{resultField, time.RFC3339Nano, dateTime.Format(time.RFC3339Nano)})
}

func (r *Result) DateString(resultField string, sampleFormat string, dateTime string) *Result {
	return r.addGenerator(DateGenerator{resultField, sampleFormat, dateTime})
}

// The BoolGenerator - used to represent simple bool value on custom FieldName()
type BoolGenerator struct {
	fieldName  string
	fieldValue bool
}

func (gen BoolGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

func (gen BoolGenerator) FieldName() string {
	return gen.fieldName
}

func (r *Result) Bool(resultField string, result bool) *Result {
	return r.addGenerator(BoolGenerator{resultField, result})
}

// The IntGenerator - used to represent simple int value on custom FieldName()
type IntGenerator struct {
	fieldName  string
	fieldValue int
}

func (gen IntGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

func (gen IntGenerator) FieldName() string {
	return gen.fieldName
}

func (r *Result) Int(resultField string, result int) *Result {
	return r.addGenerator(IntGenerator{resultField, result})
}

// The JsonValueGenerator - base struct for all Json Value Generators
type JsonValueGenerator struct {
	fieldName   string
	jqOperation string
	jsonPayload json.RawMessage
}

func (gen JsonValueGenerator) FieldName() string {
	return gen.fieldName
}

// The JsonValueStringGenerator - string result on JSON path given over JQ command
type JsonValueStringGenerator struct {
	JsonValueGenerator
}

func (gen JsonValueStringGenerator) GenerateValue() (interface{}, error) {
	parseResult, err := parseNode(gen.jqOperation, gen.jsonPayload)
	if err != nil {
		return nil, err
	}
	result, err := strconv.Unquote(string(parseResult))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Result) JsonString(resultField string, jqOp string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonValueStringGenerator{JsonValueGenerator{resultField, jqOp, jsonPayload}})
}

// The JsonValueBoolGenerator - bool result on JSON path given over JQ command
type JsonValueBoolGenerator struct {
	JsonValueGenerator
}

func (gen JsonValueBoolGenerator) GenerateValue() (interface{}, error) {
	parseResult, err := parseNode(gen.jqOperation, gen.jsonPayload)
	if err != nil {
		return nil, err
	}
	boolValue, err := strconv.ParseBool(string(parseResult))
	if err != nil {
		return nil, err
	}
	return boolValue, nil
}

func (r *Result) JsonBool(resultField string, jqOp string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonValueBoolGenerator{JsonValueGenerator{resultField, jqOp, jsonPayload}})
}

// The JsonValueDateGenerator - date result on JSON path given over JQ command
type JsonValueDateGenerator struct {
	JsonValueGenerator
	sampleFormat string
}

func (gen JsonValueDateGenerator) GenerateValue() (interface{}, error) {
	parseResult, err := parseNode(gen.jqOperation, gen.jsonPayload)
	if err != nil {
		return nil, err
	}
	result, err := strconv.Unquote(string(parseResult))
	if err != nil {
		return nil, err
	}
	date, err := parseDate(gen.sampleFormat, result)
	if err != nil {
		return nil, err
	}
	return date, nil
}

func (r *Result) JsonDate(resultField string, jqOp string, sampleFormat string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonValueDateGenerator{JsonValueGenerator{resultField, jqOp, jsonPayload}, sampleFormat})
}

// The JsonValueNodeGenerator - JSON result on JSON path given over JQ command
type JsonValueNodeGenerator struct {
	JsonValueGenerator
}

func (gen JsonValueNodeGenerator) GenerateValue() (interface{}, error) {
	parseResult, err := parseNode(gen.jqOperation, gen.jsonPayload)
	if err != nil {
		return nil, err
	}
	var commandResultMap map[string]json.RawMessage
	err = json.Unmarshal(parseResult, &commandResultMap)
	if err != nil {
		return nil, err
	}
	return commandResultMap, nil
}

func (r *Result) JsonNode(resultField string, jqOp string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonValueNodeGenerator{JsonValueGenerator{resultField, jqOp, jsonPayload}})
}

type JsonRawGenerator struct {
	fieldName   string
	jsonPayload json.RawMessage
}

func (gen JsonRawGenerator) GenerateValue() (interface{}, error) {
	return gen.jsonPayload, nil
}

func (gen JsonRawGenerator) FieldName() string {
	return gen.fieldName
}

type LookupResultElementGenerator struct {
	fieldName     string
	lookupResults []LookupResultElement
}

func (gen LookupResultElementGenerator) GenerateValue() (interface{}, error) {
	result, err := json.Marshal(gen.lookupResults)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (gen LookupResultElementGenerator) FieldName() string {
	return gen.fieldName
}

type JsonGenerator struct {
	fieldName   string
	jsonPayload json.RawMessage
}

func (gen JsonGenerator) GenerateValue() (interface{}, error) {
	var commandResultMap map[string]json.RawMessage
	err := json.Unmarshal(gen.jsonPayload, &commandResultMap)
	if err != nil {
		return nil, err
	}
	return commandResultMap, nil
}

func (gen JsonGenerator) FieldName() string {
	return gen.fieldName
}

// The StringSliceGenerator - used to represent []string value on custom FieldName()
type StringSliceGenerator struct {
	fieldName  string
	fieldValue []string
}

func (gen StringSliceGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

func (gen StringSliceGenerator) FieldName() string {
	return gen.fieldName
}

func (r *Result) StringSlice(resultField string, result []string) *Result {
	return r.addGenerator(StringSliceGenerator{resultField, result})
}

// The SliceGenerator - used to represent []interface{} value on custom FieldName()
type SliceGenerator struct {
	fieldName  string
	fieldValue []interface{}
}

func (gen SliceGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

func (gen SliceGenerator) FieldName() string {
	return gen.fieldName
}

func (r *Result) Slice(resultField string, result []interface{}) *Result {
	return r.addGenerator(SliceGenerator{resultField, result})
}

func (r *Result) Json(resultField string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonGenerator{resultField, jsonPayload})
}

func (r *Result) JsonRaw(resultField string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonRawGenerator{resultField, jsonPayload})
}

func (r *Result) LookupResultElement(resultField string, lookupResults []LookupResultElement) *Result {
	return r.addGenerator(LookupResultElementGenerator{resultField, lookupResults})
}

// CustomValue - provide custom implementation of Generator interface which will add custom parsed or generated value to the result
func (r *Result) CustomValue(generator Generator) *Result {
	return r.addGenerator(generator)
}

func (r *Result) Get() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for _, generator := range r.resultGenerators {
		value, err := generator.GenerateValue()
		// write value first if it's not empty
		if value != nil {
			result[generator.FieldName()] = value
		}
		if err != nil {
			switch generator.(type) {
			case ErrorGenerator:
				return result, err
			default:
				return result, fmt.Errorf("error evaluating '%s' field value: %v", generator.FieldName(), err)
			}
		}
	}
	return result, nil
}
