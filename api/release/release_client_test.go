package release

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"testing"
)

func TestReleaseClient(t *testing.T) {
	ConveyTest(t)
}

func ConveyTest(t *testing.T) {
	Convey("Test ReleaseApi client", t, func() {
		clientConfig := NewReleaseClient(ctx).Client.GetConfig()

		So(clientConfig.Host, ShouldEqual, "localhost:5516")
	})
}

var ctx = task.ReleaseContext{
	Id: "release-1",
	AutomatedTaskAsUser: task.AutomatedTaskAsUserContext{
		Username: "admin",
		Password: "admin",
	},
	Url: "localhost:5516",
}
