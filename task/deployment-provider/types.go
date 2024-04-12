package deployment_provider

import (
	"github.com/digital-ai/release-integration-sdk-go/task/ci"
	"reflect"
	"time"
)

func init() {
	ci.RegisterTypeMappings(ci.TypeMap{
		reflect.TypeOf(Environment{}):             "xlrelease.Environment",
		reflect.TypeOf(Application{}):             "xlrelease.Application",
		reflect.TypeOf(DeploymentState{}):         "xlrelease.DeploymentState",
		reflect.TypeOf(DeploymentProviderEvent{}): "events.DeploymentProviderEvent",
	})
}

type DeploymentProviderEvent struct {
	Operation       string          `synthetic:"operation"`
	Environment     Environment     `synthetic:"environment"`
	Application     Application     `synthetic:"application"`
	DeploymentState DeploymentState `synthetic:"deploymentState"`
}

type Environment struct {
	Title            string           `synthetic:"title"`
	CorrelationUid   string           `synthetic:"correlationUid"`
	DeploymentTarget DeploymentTarget `synthetic:"deploymentTarget"`
}

type Application struct {
	Title             string            `synthetic:"title"`
	CorrelationUid    string            `synthetic:"correlationUid"`
	ApplicationSource ApplicationSource `synthetic:"applicationSource"`
}

type DeploymentState struct {
	StatusGroup    string    `synthetic:"statusGroup"`
	Status         string    `synthetic:"status"`
	VersionTag     string    `synthetic:"versionTag"`
	VersionState   string    `synthetic:"versionState"`
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
