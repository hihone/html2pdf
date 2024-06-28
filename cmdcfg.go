//go:build !windows
// +build !windows

package html2pdf

import "os/exec"

func cmdConfig(cmd *exec.Cmd) {}
