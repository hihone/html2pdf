//go:build windows
// +build windows

package html2pdf

import (
	"os/exec"
	"syscall"
)

func cmdConfig(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000, HideWindow: true}
}
