package vmgr

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"rvhc/daemon/internal/dataimg"
	"rvhc/daemon/internal/db"
	"rvhc/downloader"
)

func New(vm db.VM) *exec.Cmd {
	return exec.Command(
		"qemu-system-riscv64",
		"-M", "virt",
		"-bios", filepath.Join(downloader.BaseDir, "bios.bin"),
		"-kernel", filepath.Join(downloader.BaseDir, "kernel.bin"),
		"-append", "rootwait root=/dev/vda ro",
		"-drive", fmt.Sprintf(
			"file=%s,format=raw,if=virtio,read-only=on",
			filepath.Join(downloader.BaseDir, "sys.img"),
		),
		"-drive", fmt.Sprintf(
			"file=%s,format=raw,if=virtio",
			filepath.Join(dataimg.BaseDir, vm.ID+".img"),
		),
		"-netdev", "user,id=net0",
		"-device", "virtio-net-device,netdev=net0",
		"-display", "none",
		"-serial", "none",
		"-chardev", "stdio,id=mon0",
		"-mon", "chardev=mon0,mode=control",
		"-smp", fmt.Sprintf("%d", vm.CPU),
		"-m", fmt.Sprintf("%dM", vm.Memory),
	)
}
