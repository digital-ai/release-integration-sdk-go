package task

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"k8s.io/klog/v2"
	"net/http"
	"os"
	"sync"
)

func HandleSuccess(result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
	}
	handleResult(outputContext)
}

func HandleError(err error, result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         -1,
		OutputProperties: result,
		JobErrorMessage:  err.Error(),
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
	wg.Add(2)
	go func() {
		defer wg.Done()
		remoteRunnerCallback(encryptedData)
	}()
	go func() {
		defer wg.Done()
		writeOutput(encryptedData)
	}()
	wg.Wait()
}

func remoteRunnerCallback(encryptedData []byte) {
	encodedCallBackUrl := os.Getenv("CALLBACK_URL")
	if len(encodedCallBackUrl) > 0 {
		callBackUrl, err := base64.StdEncoding.DecodeString(encodedCallBackUrl)
		if err != nil {
			klog.Errorf("Cannot decode Callback URL %s, skipping - output written to output file", err)
		}
		url := string(callBackUrl)
		// TODO retry schema maybe?
		response, httpError := http.Post(url, "application/json", bytes.NewReader(encryptedData))
		if httpError != nil {
			klog.Errorf("Cannot finish Callback request: %s, skipping - output written to output file", httpError)
		}
		if response.StatusCode != 200 {
			klog.Errorf("Got NOK HTTP Status Code %s, skipping - output written to output file", response.Status)
		}
	}
}
