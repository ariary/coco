package main

import (
	"fmt"
	"log"

	"github.com/ariary/coco/pkg/c2c"
	encryption "github.com/ariary/go-utils/pkg/encrypt"
	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	db := c2c.ModuleDB{}
	socketName := c2c.SOCKET + "." + encryption.GenerateRandom()
	sc, err := ipc.StartServer(socketName, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Start ipc server for socket:", socketName)

	c2c.WaitConnection(sc, &db)

	//send instruction
	// c2c.CheckSendMessage(sc, "world")
	// time.Sleep(3 * time.Second)
	// c2c.CheckSendMessage(sc, "toto")
	// time.Sleep(3 * time.Second)
	resp := c2c.CheckSendMessageAndWaitResponse(sc, "youhou")
	fmt.Println(resp)
}

//goroutine wait response?
