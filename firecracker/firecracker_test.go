package firecracker

import (
	"context"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func testPath() string  {
	gopath, ok := os.LookupEnv("GOPATH")
	So(ok, ShouldBeTrue)

	git_prov, ok := os.LookupEnv("GIT_PROVIDER")
	So(ok, ShouldBeTrue)

	git_repo, ok := os.LookupEnv("GIT_REPO")
	So(ok, ShouldBeTrue)

	return path.Join(gopath, "src", git_prov, git_repo, ".test")
}

func socketPath(socket string) string {
	ex, err := os.Executable()
	So(err, ShouldBeNil)
	execPath := filepath.Dir(ex)
	return path.Join(execPath, socket)
}

func setupTestVM(socket string) {
	tp := testPath()
	st, err := os.Lstat(tp)
	So(err, ShouldBeNil)
	So(st.IsDir(), ShouldBeTrue)

	client, err := New(socketPath(socket))
	So(err, ShouldBeNil)
	So(client, ShouldNotBeNil)

	err = client.SetBootSource(path.Join(tp, "hello-vmlinux.bin"), "console=ttyS0 reboot=k panic=1 pci=off")
	So(err, ShouldBeNil)

	err = client.SetDrive(&StorageDrive{
		ID:         "rootfs",
		HostPath:   path.Join(tp, "hello-rootfs.ext4"),
		IsRoot:     true,
		IsReadOnly: false,
	})
	So(err, ShouldBeNil)

	err = client.UpdateDrivePath("rootfs", path.Join(tp, "hello-rootfs.ext4"))
	So(err, ShouldBeNil)

	conf := &MachineConfig{
		CPUTemplate:C3Template,
		VCpuCount: 1,
		MemSizeMb: 32,
	}
	err = client.SetConfig(conf)
	So(err, ShouldBeNil)

	err = client.Start()
	So(err, ShouldBeNil)

	state, err := client.State()
	So(err, ShouldBeNil)
	So(state, ShouldEqual, Running)

	readConfig, err := client.Config()
	So(err, ShouldBeNil)
	So(readConfig, ShouldResemble, conf)

	err = client.Rescan("rootfs")
	So(err, ShouldBeNil)
}

func startFirecracker(socket string) *exec.Cmd {
	sockp := socketPath(socket)
	cmd := exec.Command("firecracker", "--api-sock", sockp)
	err := cmd.Start()
	So(err, ShouldBeNil)

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 5)
	func(ctx context.Context) {
		for {
			_, err := os.Lstat(sockp)
			if err == nil {
				return
			}

			select {
			case <-ctx.Done():
				So("Timeout", ShouldBeTrue)
			default:
				continue
			}
		}
	}(ctx)

	_, err = os.Lstat(sockp)
	So(err, ShouldBeNil)

	return cmd
}

func stopFirecracker(cracker *exec.Cmd) {
	err := cracker.Process.Signal(syscall.SIGTERM)
	So(err, ShouldBeNil)
}

func TestFirecracker(t *testing.T) {
	Convey("Given a test Firecracker VM", t, func() {
		Convey("Should be able to get instance ID and Status", func() {
			socket := "StateTest.sock"
			crackerProc := startFirecracker(socket)
			defer stopFirecracker(crackerProc)

			client, err := New(socketPath(socket))
			So(err, ShouldBeNil)
			So(client, ShouldNotBeNil)

			id, err := client.ID()
			So(err, ShouldBeNil)
			So(id, ShouldNotBeEmpty)

			So(id, ShouldEqual, "anonymous-instance")

			state, err := client.State()
			So(err, ShouldBeNil)
			So(state, ShouldNotBeEmpty)

		})

		Convey("Should be able to create a bootable instance", func() {
			socket := "SetupTest.sock"
			crackerProc := startFirecracker(socket)
			defer stopFirecracker(crackerProc)

			setupTestVM(socket)
		})
	})
}
