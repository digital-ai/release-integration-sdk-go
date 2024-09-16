package deployment_provider

import (
	"github.com/digital-ai/release-integration-sdk-go/task/ci"
	"reflect"
	"time"
)

func init() {
	ci.RegisterTypeMappings(ci.TypeMap{
		reflect.TypeOf(DeploymentState{}):         "xlrelease.DeploymentState",
		reflect.TypeOf(DeploymentProviderEvent{}): "events.DeploymentProviderEvent",
	})
}

type DeploymentProviderEvent struct {
	Operation         string            `synthetic:"operation"`
	ApplicationCiud   string            `synthetic:"applicationCiud"`
	ApplicationTitle  string            `synthetic:"applicationTitle"`
	ApplicationSource ApplicationSource `synthetic:"applicationSource"`
	EnvironmentCiud   string            `synthetic:"environmentCiud"`
	EnvironmentTitle  string            `synthetic:"environmentTitle"`
	DeploymentTarget  DeploymentTarget  `synthetic:"deploymentTarget"`
	DeploymentState   DeploymentState   `synthetic:"deploymentState"`
	ConfigId          string            `synthetic:"configId"`
}

type DeploymentState struct {
	StatusGroup    string    `synthetic:"statusGroup"`
	Status         string    `synthetic:"status"`
	DeploymentLink string    `synthetic:"deploymentLink"`
	VersionTag     string    `synthetic:"versionTag"`
	VersionState   string    `synthetic:"versionState"`
	VersionTooltip string    `synthetic:"versionTooltip"`
	VersionLink    string    `synthetic:"versionLink"`
	DeploymentType string    `synthetic:"deploymentType"`
	User           string    `synthetic:"user"`
	LastChangeTime time.Time `synthetic:"lastChangeTime"`
}

type ApplicationSourceType struct {
	ServerUrl string `synthetic:"serverUrl"`
}

type DeploymentTargetType struct {
	TargetUrl string `synthetic:"targetUrl"`
}

type ComposedType interface {
	Type() reflect.Type
}

type ApplicationSource interface {
	ComposedType
}

type DeploymentTarget interface {
	ComposedType
}

func RegisterApplicationSourceType(typeName string) {
	ci.RegisterTypeMapping(reflect.TypeOf(ApplicationSourceType{}), typeName)
}

func RegisterDeploymentTargetType(typeName string) {
	ci.RegisterTypeMapping(reflect.TypeOf(DeploymentTargetType{}), typeName)
}
