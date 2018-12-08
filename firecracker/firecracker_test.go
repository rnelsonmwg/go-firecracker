package firecracker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFirecracker(t *testing.T) {
	Convey("Given test account", t, func() {
		So(testEmail, ShouldNotBeEmpty)
		So(testPassword, ShouldNotBeEmpty)

		toggl := SignIn(testEmail, testPassword)

		shouldBeReused := func(name string) (w *Workspace) {
			var err error
			w, err = toggl.WorkspaceByName(name)
			if err != nil {
				w, err = toggl.NewWorkspace(name)
				So(err, ShouldBeNil)
				So(w, ShouldNotBeNil)
				So(w.Name, ShouldEqual, name)
			}

			return w
		}

		Convey("Should be able to to create a temporary workspace", func() {
			w := shouldBeReused("temp_workspace")

			Convey("Should be able to read it back", func() {
				var err error

				w, err = toggl.WorkspaceByName("temp_workspace")
				So(err, ShouldBeNil)
				So(w, ShouldNotBeNil)
				So(w.Name, ShouldEqual, "temp_workspace")

				users, err := toggl.WorkspaceUsers("temp_workspace")
				So(err, ShouldBeNil)
				So(users, ShouldNotBeEmpty)

				Convey("and delete afterwards", func() {
					err := toggl.LeaveWorkspace("temp_workspace")
					So(err, ShouldBeNil)

					w, err := toggl.WorkspaceByName("temp_workspace")
					So(w, ShouldBeNil)
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("Should be able to rename a temporary workspace", func() {
			var err error
			w := shouldBeReused("temp_workspace_renamed")

			w.Name = "temp_workspace_renamed"
			err = toggl.UpdateWorkspace(w)
			So(err, ShouldBeNil)

			w, err = toggl.WorkspaceByName("temp_workspace_renamed")
			So(err, ShouldBeNil)
			So(w, ShouldNotBeNil)
			So(w.Name, ShouldEqual, "temp_workspace_renamed")

			Convey("and delete afterwards", func() {
				err := toggl.LeaveWorkspace("temp_workspace_renamed")
				So(err, ShouldBeNil)

				w, err := toggl.WorkspaceByName("temp_workspace_renamed")
				So(w, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
