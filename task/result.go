package task

import (
	"encoding/json"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/util"
	"strconv"
	"time"
)

const (
	errorMessage = "errorMessage"
)

type AbortError struct{}

func (e *AbortError) Error() string {
	return "Task was aborted"
}

// Result represents the result of a task execution.
type Result struct {
	resultGenerators []Generator
	reportingRecords []interface{}
}

// NewResult creates a new Result instance.
func NewResult() *Result {
	return &Result{
		resultGenerators: []Generator{},
		reportingRecords: []interface{}{},
	}
}

// Generator is an interface for result value generators.
type Generator interface {
	GenerateValue() (interface{}, error)
	FieldName() string
}

// addGenerator adds a generator to the result.
func (r *Result) addGenerator(generator Generator) *Result {
	r.resultGenerators = append(r.resultGenerators, generator)
	return r
}

// The AbortGenerator - used to represent aborted execution result
type AbortGenerator struct {
	result map[string]interface{}
}

func (gen AbortGenerator) GenerateValue() (interface{}, error) {
	return "", &AbortError{}
}

func (gen AbortGenerator) FieldName() string {
	return ""
}

func (r *Result) Aborted(result *Result) *Result {
	resultMap, err := result.Get()
	if err != nil {
		// add empty result
		return r.addGenerator(AbortGenerator{})
	}
	return r.addGenerator(AbortGenerator{result: resultMap})
}

// addReportingRecord adds a record to the result.
func (r *Result) addReportingRecord(record interface{}) *Result {
	r.reportingRecords = append(r.reportingRecords, record)
	return r
}

// ErrorGenerator represents an error in the standardized response.
type ErrorGenerator struct {
	err error
}

// GenerateValue generates the error value.
func (gen ErrorGenerator) GenerateValue() (interface{}, error) {
	return gen.err.Error(), gen.err
}

// FieldName returns the field name for the error.
func (gen ErrorGenerator) FieldName() string {
	return errorMessage
}

// Error adds an error to the result.
func (r *Result) Error(err error) *Result {
	return r.addGenerator(ErrorGenerator{err})
}

// StringGenerator represents a simple string value with a custom field name.
type StringGenerator struct {
	fieldName  string
	fieldValue string
}

// GenerateValue generates the string value.
func (gen StringGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

// FieldName returns the custom field name for the string value.
func (gen StringGenerator) FieldName() string {
	return gen.fieldName
}

// String adds a string value with a custom field name to the result.
func (r *Result) String(resultField string, result string) *Result {
	return r.addGenerator(StringGenerator{resultField, result})
}

// DateGenerator represents a simple date and time value with a custom field name.
type DateGenerator struct {
	fieldName    string
	sampleFormat string
	dateTime     string
}

// GenerateValue generates the date and time value.
func (gen DateGenerator) GenerateValue() (interface{}, error) {
	date, err := util.ParseDate(gen.sampleFormat, gen.dateTime)
	if err != nil {
		return nil, err
	}
	return date, nil
}

// FieldName returns the custom field name for the date and time value.
func (gen DateGenerator) FieldName() string {
	return gen.fieldName
}

// Date adds a date and time value with a custom field name to the result.
func (r *Result) Date(resultField string, dateTime time.Time) *Result {
	return r.addGenerator(DateGenerator{resultField, time.RFC3339Nano, dateTime.Format(time.RFC3339Nano)})
}

// DateString adds a date and time value with a custom format and field name to the result.
func (r *Result) DateString(resultField string, sampleFormat string, dateTime string) *Result {
	return r.addGenerator(DateGenerator{resultField, sampleFormat, dateTime})
}

// BoolGenerator represents a simple boolean value with a custom field name.
type BoolGenerator struct {
	fieldName  string
	fieldValue bool
}

// GenerateValue generates the boolean value.
func (gen BoolGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

// FieldName returns the custom field name for the boolean value.
func (gen BoolGenerator) FieldName() string {
	return gen.fieldName
}

// Bool adds a boolean value with a custom field name to the result.
func (r *Result) Bool(resultField string, result bool) *Result {
	return r.addGenerator(BoolGenerator{resultField, result})
}

// IntGenerator represents a simple integer value with a custom field name.
type IntGenerator struct {
	fieldName  string
	fieldValue int
}

// GenerateValue generates the integer value.
func (gen IntGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

// FieldName returns the custom field name for the integer value.
func (gen IntGenerator) FieldName() string {
	return gen.fieldName
}

// Int adds an integer value with a custom field name to the result.
func (r *Result) Int(resultField string, result int) *Result {
	return r.addGenerator(IntGenerator{resultField, result})
}

// JsonValueGenerator is the base struct for all JSON value generators.
type JsonValueGenerator struct {
	fieldName   string
	jqOperation string
	jsonPayload json.RawMessage
}

// FieldName returns the field name for the JSON value generator.
func (gen JsonValueGenerator) FieldName() string {
	return gen.fieldName
}

// JsonValueStringGenerator represents a string result on a JSON path given over a JQ command.
type JsonValueStringGenerator struct {
	JsonValueGenerator
}

// GenerateValue generates the string value.
func (gen JsonValueStringGenerator) GenerateValue() (interface{}, error) {
	parseResult, err := util.ParseNode(gen.jqOperation, gen.jsonPayload)
	if err != nil {
		return nil, err
	}
	result, err := strconv.Unquote(string(parseResult))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// JsonString adds a string result on a JSON path given over a JQ command to the result.
func (r *Result) JsonString(resultField string, jqOp string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonValueStringGenerator{JsonValueGenerator{resultField, jqOp, jsonPayload}})
}

// JsonValueBoolGenerator represents a boolean result on a JSON path given over a JQ command.
type JsonValueBoolGenerator struct {
	JsonValueGenerator
}

// GenerateValue generates the boolean value.
func (gen JsonValueBoolGenerator) GenerateValue() (interface{}, error) {
	parseResult, err := util.ParseNode(gen.jqOperation, gen.jsonPayload)
	if err != nil {
		return nil, err
	}
	boolValue, err := strconv.ParseBool(string(parseResult))
	if err != nil {
		return nil, err
	}
	return boolValue, nil
}

// JsonBool adds a boolean result on a JSON path given over a JQ command to the result.
func (r *Result) JsonBool(resultField string, jqOp string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonValueBoolGenerator{JsonValueGenerator{resultField, jqOp, jsonPayload}})
}

// JsonValueDateGenerator represents a date result on a JSON path given over a JQ command.
type JsonValueDateGenerator struct {
	JsonValueGenerator
	sampleFormat string
}

// GenerateValue generates the date value.
func (gen JsonValueDateGenerator) GenerateValue() (interface{}, error) {
	parseResult, err := util.ParseNode(gen.jqOperation, gen.jsonPayload)
	if err != nil {
		return nil, err
	}
	result, err := strconv.Unquote(string(parseResult))
	if err != nil {
		return nil, err
	}
	date, err := util.ParseDate(gen.sampleFormat, result)
	if err != nil {
		return nil, err
	}
	return date, nil
}

// JsonDate adds a date result on a JSON path given over a JQ command to the result.
func (r *Result) JsonDate(resultField string, jqOp string, sampleFormat string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonValueDateGenerator{JsonValueGenerator{resultField, jqOp, jsonPayload}, sampleFormat})
}

// JsonValueNodeGenerator represents a JSON result on a JSON path given over a JQ command.
type JsonValueNodeGenerator struct {
	JsonValueGenerator
}

// GenerateValue generates the JSON value.
func (gen JsonValueNodeGenerator) GenerateValue() (interface{}, error) {
	parseResult, err := util.ParseNode(gen.jqOperation, gen.jsonPayload)
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

// JsonNode adds a JSON result on a JSON path given over a JQ command to the result.
func (r *Result) JsonNode(resultField string, jqOp string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonValueNodeGenerator{JsonValueGenerator{resultField, jqOp, jsonPayload}})
}

// JsonRawGenerator represents a raw JSON value with a custom field name.
type JsonRawGenerator struct {
	fieldName   string
	jsonPayload json.RawMessage
}

// GenerateValue generates the raw JSON value.
func (gen JsonRawGenerator) GenerateValue() (interface{}, error) {
	return gen.jsonPayload, nil
}

// FieldName returns the custom field name for the raw JSON value.
func (gen JsonRawGenerator) FieldName() string {
	return gen.fieldName
}

// LookupResultElementsGenerator represents lookup result elements with a custom field name.
type LookupResultElementsGenerator struct {
	fieldName     string
	lookupResults []LookupResultElement
}

// GenerateValue generates the lookup result elements value.
func (gen LookupResultElementsGenerator) GenerateValue() (interface{}, error) {
	result, err := json.Marshal(gen.lookupResults)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(result), nil
}

// FieldName returns the custom field name for the lookup result elements value.
func (gen LookupResultElementsGenerator) FieldName() string {
	return gen.fieldName
}

// JsonGenerator represents a JSON value with a custom field name.
type JsonGenerator struct {
	fieldName   string
	jsonPayload json.RawMessage
}

// GenerateValue generates the JSON value.
func (gen JsonGenerator) GenerateValue() (interface{}, error) {
	var commandResultMap map[string]json.RawMessage
	err := json.Unmarshal(gen.jsonPayload, &commandResultMap)
	if err != nil {
		return nil, err
	}
	return commandResultMap, nil
}

// FieldName returns the custom field name for the JSON value.
func (gen JsonGenerator) FieldName() string {
	return gen.fieldName
}

// StringSliceGenerator represents a []string value with a custom field name.
type StringSliceGenerator struct {
	fieldName  string
	fieldValue []string
}

// GenerateValue generates the []string value.
func (gen StringSliceGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

// FieldName returns the custom field name for the []string value.
func (gen StringSliceGenerator) FieldName() string {
	return gen.fieldName
}

// StringSlice adds a []string value with a custom field name to the result.
func (r *Result) StringSlice(resultField string, result []string) *Result {
	return r.addGenerator(StringSliceGenerator{resultField, result})
}

// SliceGenerator represents a []interface{} value with a custom field name.
type SliceGenerator struct {
	fieldName  string
	fieldValue []interface{}
}

// GenerateValue generates the []interface{} value.
func (gen SliceGenerator) GenerateValue() (interface{}, error) {
	return gen.fieldValue, nil
}

// FieldName returns the custom field name for the []interface{} value.
func (gen SliceGenerator) FieldName() string {
	return gen.fieldName
}

// Slice adds a []interface{} value with a custom field name to the result.
func (r *Result) Slice(resultField string, result []interface{}) *Result {
	return r.addGenerator(SliceGenerator{resultField, result})
}

// Json adds a JSON value with a custom field name to the result.
func (r *Result) Json(resultField string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonGenerator{resultField, jsonPayload})
}

// JsonRaw adds a raw JSON value with a custom field name to the result.
func (r *Result) JsonRaw(resultField string, jsonPayload json.RawMessage) *Result {
	return r.addGenerator(JsonRawGenerator{resultField, jsonPayload})
}

// LookupResultElements adds lookup result elements with a custom field name to the result.
func (r *Result) LookupResultElements(resultField string, lookupResults []LookupResultElement) *Result {
	return r.addGenerator(LookupResultElementsGenerator{resultField, lookupResults})
}

// CustomValue provides a custom implementation of the Generator interface, adding a custom parsed or generated value to the result.
func (r *Result) CustomValue(generator Generator) *Result {
	return r.addGenerator(generator)
}

// ReportingRecord adds reporting record to the result.
func (r *Result) ReportingRecord(record interface{}) *Result {
	return r.addReportingRecord(record)
}

// CreateDeploymentRecord creates a new reporting record for deployment.
func CreateDeploymentRecord(client *http.HttpClient, taskInfo *DeploymentRecordTaskInfo, status DeploymentStatus) *DeploymentRecord {
	serverUrl := client.GetBaseUrl()
	return NewReportingRecord().
		TargetId(taskInfo.ReleaseTaskId).
		ServerUser(taskInfo.Username).
		ServerUrl(serverUrl).
		DeploymentRecord().
		TaskId(taskInfo.TaskId).
		DeploymentStatus(status).
		DeploymentPackage(taskInfo.DeploymentPackage).
		DeployedApplication(taskInfo.DeployedApplication).
		ApplicationVersion(taskInfo.DeploymentVersion).
		DeploymentEnvironment(taskInfo.DeploymentEnvironment).
		Build()
}

// Get returns map with all values that were added to result
func (r *Result) Get() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for _, generator := range r.resultGenerators {
		value, err := generator.GenerateValue()
		// write value first if it's not empty
		if value != nil {
			result[generator.FieldName()] = value
		}
		if err != nil {
			switch g := generator.(type) {
			case AbortGenerator:
				return g.result, err
			case ErrorGenerator:
				return result, err
			default:
				return result, fmt.Errorf("error evaluating '%s' field value: %v", generator.FieldName(), err)
			}
		}
	}
	return result, nil
}

// GetRecords returns list of all reporting records
func (r *Result) GetRecords() []interface{} {
	return r.reportingRecords
}
