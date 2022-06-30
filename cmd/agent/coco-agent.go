package main

import (
	"fmt"
	"log"
	"time"

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
	// resp := c2c.CheckSendMessageAndWaitResponse(sc, "youhou")
	// fmt.Println(resp)
	sayHelloInstr := c2c.Instruction{Type: c2c.Run}
	if err := c2c.SendInstruction(sc, sayHelloInstr); err != nil {
		fmt.Println("failed sending instruciton:", err)
	}
	if resp, err := c2c.WaitResponse(sc); err != nil {
		fmt.Println("Error while waiting for response:", err)
	} else {
		fmt.Println(resp)
	}
	time.Sleep(2 * time.Second)
	//kill
	c2c.KillModule(sc)

}

//goroutine wait response?
