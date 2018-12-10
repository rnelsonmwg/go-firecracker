package firecracker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var testFirecrackerSock = "/firecracker/firecracker.sock"

func TestFirecracker(t *testing.T) {
	Convey("Given a test Firecracker instance", t, func() {
		So(testFirecrackerSock, ShouldNotBeEmpty)

		Convey("Should be able to get instance ID and Status", func() {
			client := NewSocket(testFirecrackerSock)
			So(client, ShouldNotBeNil)

			id, err := client.ID()
			So(err, ShouldBeNil)
			So(id, ShouldNotBeEmpty)

			state, err := client.State()
			So(err, ShouldBeNil)
			So(state, ShouldNotBeEmpty)
		})
	})
}
