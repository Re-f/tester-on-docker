// +build !windows

package docker

import (
	"fmt"
	"os/exec"
	"strings"
)

func executeOnDocker(str string) (string, error) {
	return execute(fmt.Sprintf("%v %v", getSudo(), str))
}

func newCmd(cmd string) *exec.Cmd {
	return exec.Command("sh", "-c", cmd)
}

func getCrossCompileCmd(pkName, os, arch string) string {
	cmds := []string{
		"CGO_ENABLED=0",
		"GOOS=" + os,
		"GOARCH=" + arch,
		"go test -c -tags inner " + pkName,
	}
	return strings.Join(cmds, " ")
}
