package test

import (
	"encoding/json"
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/runner"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"time"
)

var jsonPayload json.RawMessage = []byte(`{"metadata":{
											  "name":"kubernetes",
											  "namespace":"default",
											  "uid":"f1b86cf1-c600-428a-9aa3-984effb91bc0",
											  "resourceVersion":"197",
											  "creationTimestamp":"2022-09-16T11:35:02Z",
											  "labels":{
												 "component":"apiserver",
												 "provider":"kubernetes"
											  },
											  "managedFields":[
												 {
													"manager":"k3s",
													"operation":"Update",
													"apiVersion":"v1",
													"time":"2022-09-16T11:35:02Z",
													"fieldsType":"FieldsV1"
												 }
											  ]
										   },
										   "spec":{
											  "ports":[
												 {
													"name":"https",
													"protocol":"TCP",
													"port":443,
													"targetPort":6443
												 }
											  ],
											  "clusterIP":"10.43.0.1",
											  "clusterIPs":[
												 "10.43.0.1"
											  ],
											  "type":"ClusterIP",
											  "sessionAffinity":"None",
											  "ipFamilies":[
												 "IPv4"
											  ],
											  "ipFamilyPolicy":"SingleStack",
											  "internalTrafficPolicy":"Cluster"
										   },
										   "just-random-bool": true
										}`)

func simpleRunnerWith(result *task.Result) runner.Runner {
	return runner.NewSimpleRunner(
		func(_ task.InputContext) *task.Result {
			return result
		})
}

var independentOutputRunner = runner.NewSimpleRunner(
	func(_ task.InputContext) *task.Result {
		return task.NewResult().String("test", "the-result")
	})

var failedTestRunner = runner.NewSimpleRunner(
	func(_ task.InputContext) *task.Result {
		return task.NewResult().Error(fmt.Errorf("this simulates an error"))
	})

var stringPropertyRunner = runner.NewSimpleRunner(
	func(input task.InputContext) *task.Result {
		return task.NewResult().String("the-type", input.Task.Type)
	})

var datePropertyRunner = runner.NewSimpleRunner(
	func(input task.InputContext) *task.Result {
		theTime := time.Date(2023, 1, 11, 10, 44, 45, 142, time.Local)
		return task.NewResult().Date("the-date", theTime)
	})

var composedPropertyRunner = runner.NewSimpleRunner(
	func(input task.InputContext) *task.Result {
		theTime := time.Date(2023, 1, 11, 10, 44, 45, 142, time.Local)
		return task.NewResult().
			Date("the-date", theTime).
			String("the-type", input.Task.Type).
			JsonString("the-json-string", ".metadata.name", jsonPayload)
	})
