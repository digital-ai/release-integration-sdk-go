package deployment_server

import (
	"encoding/json"
	"github.com/digital-ai/release-integration-sdk-go/task/ci"
	"reflect"
	"testing"
	"time"
)

type PluginDeploymentTarget struct {
	DeploymentTargetType
	EnvironmentPath string `synthetic:"environmentPath"`
}

func (das PluginDeploymentTarget) Type() reflect.Type {
	return reflect.ValueOf(PluginDeploymentTarget{}).Type()
}

type PluginApplicationSource struct {
	ApplicationSourceType
	ApplicationPath string `synthetic:"applicationPath"`
	ApplicationType string `synthetic:"applicationType"`
}

func (das PluginApplicationSource) Type() reflect.Type {
	return reflect.TypeOf(PluginApplicationSource{})
}

func TestSerializeCi(t *testing.T) {
	type Tests struct {
		name                 string
		event                DeploymentServerEvent
		appSourceTypeName    string
		deployTargetTypeName string
		returnValue          map[string]interface{}
		expectedErr          error
	}

	timeDate, _ := time.Parse(time.RFC3339, "2024-04-20T05:35:00+01:00")
	tests := []Tests{
		{
			name: "status-filter",
			event: DeploymentServerEvent{
				Operation: "create",

				ApplicationTitle: "important-app",
				ApplicationCuid:  "UUID-APP",
				ApplicationSource: PluginApplicationSource{
					ApplicationSourceType: ApplicationSourceType{
						ServerUrl: "test-server-url",
					},
					ApplicationType: "Ear",
					ApplicationPath: "Applications/important-app",
				},
				DeploymentState: DeploymentState{
					Status:         "crashed",
					StatusGroup:    "failed",
					DeploymentType: "rollback",
					LastChangeTime: timeDate,
					User:           "rollback-user",
					VersionTag:     "1.2",
				},

				EnvironmentTitle: "production-important",
				EnvironmentCuid:  "UUID-ENV",
				DeploymentTarget: PluginDeploymentTarget{
					DeploymentTargetType: DeploymentTargetType{
						TargetUrl: "target-url",
					},
					EnvironmentPath: "Environments/production",
				},
			},
			appSourceTypeName:    "plugin.ApplicationSource",
			deployTargetTypeName: "plugin.DeploymentTarget",
			returnValue: map[string]interface{}{
				"operation": "create",
				"id":        nil,
				"type":      "events.DeploymentServerEvent",

				"applicationTitle": "important-app",
				"applicationCuid":  "UUID-APP",
				"applicationSource": map[string]interface{}{
					"id":              nil,
					"type":            "plugin.ApplicationSource",
					"applicationPath": "Applications/important-app",
					"applicationType": "Ear",
					"serverUrl":       "test-server-url",
				},
				"deploymentState": map[string]interface{}{
					"id":             nil,
					"type":           "xlrelease.DeploymentState",
					"status":         "crashed",
					"statusGroup":    "failed",
					"deploymentType": "rollback",
					"lastChangeTime": "2024-04-20T05:35:00+01:00",
					"user":           "rollback-user",
					"versionTag":     "1.2",
				},

				"environmentTitle": "production-important",
				"environmentCuid":  "UUID-ENV",
				"deploymentTarget": map[string]interface{}{
					"id":              nil,
					"type":            "plugin.DeploymentTarget",
					"targetUrl":       "target-url",
					"environmentPath": "Environments/production",
				},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterApplicationSourceType(tt.appSourceTypeName)
			RegisterDeploymentTargetType(tt.deployTargetTypeName)
			got, err := ci.SerializeCi(tt.event)
			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Fatalf("Actual: `%v`; Expected: `%v`", err, tt.expectedErr)
			}
			var data map[string]interface{}
			err = json.Unmarshal(got, &data)
			if err != nil {
				t.Fatalf("Actual: `%v` - couln't unmarshal to map", string(got))
			}
			if !reflect.DeepEqual(data, tt.returnValue) {
				t.Fatalf("Actual: `%v`; Expected: `%v`", data, tt.returnValue)
			} else {
				t.Logf("Success!")
			}
		})
	}
}
