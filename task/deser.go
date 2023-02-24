package task

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/klog/v2"
	"net/http"
	"os"
	"sync"
)

const (
	InputCategory  = "input"
	InputLocation  = "INPUT_LOCATION"
	OutputLocation = "OUTPUT_LOCATION"
)

func Deserialize(context *InputContext) error {
	var inputLocation = os.Getenv(InputLocation)
	inputContent, err := os.Open(inputLocation)
	// if we os.Open returns an error then handle it
	if err != nil {
		klog.Errorf("Cannot open: '%s' [%v]", inputLocation, err)
		return err
	}
	// defer the closing of our inputContent so that we can parse it later on
	defer inputContent.Close()

	content, err := io.ReadAll(inputContent)
	if err != nil {
		return err
	}
	decrypted, err := Decrypt(content)
	if err != nil {
		return err
	}

	decoded, err := base64.StdEncoding.DecodeString(string(decrypted))
	if err != nil {
		return err
	}
	unMarshalErr := json.Unmarshal(decoded, context)
	if unMarshalErr != nil {
		klog.Errorf("Cannot umarshal input: %v", unMarshalErr)
		return unMarshalErr
	}

	return nil
}

func DeserializeTask(properties []PropertyDefinition, taskInstance any) error {
	var inputs []PropertyDefinition
	for _, property := range properties {
		if property.Category == InputCategory {
			inputs = append(inputs, property)
		}
	}

	return UnmarshalProperties(inputs, taskInstance)
}

func Serialize(result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
	}
	handleResult(outputContext)
}

func SerializeError(result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         -1,
		OutputProperties: result,
	}
	handleResult(outputContext)
}

func handleResult(outputContext TaskOutputContext) {
	data, _ := json.Marshal(outputContext)
	encryptedData, encryptErr := Encrypt(data)
	if encryptErr != nil {
		klog.Fatalf("error encrypting data [%v]", encryptErr)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		remoteRunnerCallback(encryptedData, outputContext.ExitCode)
	}()
	fmt.Println("Waiting for goroutine to finish")
	wg.Wait()

	writeOutput(encryptedData)
}

func remoteRunnerCallback(encryptedData []byte, exitCode int64) {
	encodedCallBackUrl := os.Getenv("CALLBACK_URL")
	if len(encodedCallBackUrl) > 0 {
		callBackUrl, err := base64.StdEncoding.DecodeString(encodedCallBackUrl)
		if err != nil {
			klog.Errorf("Cannot decode Callback URL %s, skipping - output written to output file", err)
		}
		url := string(callBackUrl)
		if exitCode != 0 {
			url = fmt.Sprintf("%s/failed", url)
		}
		// TODO retry schema maybe?
		_, httpError := http.Post(url, "application/json", bytes.NewReader(encryptedData))
		if httpError != nil {
			klog.Errorf("Cannot finish Callback request: %s, skipping - output written to output file", httpError)
		}
	}
}

func writeOutput(encryptedData []byte) {
	outputLocation := os.Getenv(OutputLocation)
	err := os.WriteFile(outputLocation, encryptedData, 0644)
	if err != nil {
		klog.Fatalf("Cannot write output to: %s [%v]", outputLocation, err)
	}
}

func UnmarshalProperties(properties []PropertyDefinition, prototype interface{}) error {
	propsMap := make(map[string]json.RawMessage)
	for _, property := range properties {
		propsMap[property.Name] = property.Value
	}
	jsonMap, err := json.Marshal(propsMap)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonMap, prototype); err != nil {
		return err
	}
	return nil
}
