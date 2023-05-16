package release

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReleaseClient(t *testing.T) {
	ConveyTest(t)
}

func ConveyTest(t *testing.T) {
	Convey("Test ReleaseApi client", t, func() {
		apiClient, err := NewReleaseApiClient(ctx)
		if err != nil {
			t.Fail()
		}

		clientConfig := apiClient.GetConfig()
		So(clientConfig.Host, ShouldEqual, "localhost:5516")
	})
}

var ctx = task.ReleaseContext{
	Id: "release-1",
	AutomatedTaskAsUser: task.AutomatedTaskAsUserContext{
		Username: "admin",
		Password: "admin",
	},
	Url: "http://localhost:5516",
}
