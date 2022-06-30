package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ariary/coco/pkg/agent"
	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	var url string
	flag.StringVar(&url, "u", "", "url testing purpose")
	flag.Parse()

	db := agent.ModuleDB{}
	module, err := agent.GetModuleContentHTTP(url)
	if err != nil {
		fmt.Println("failed to retrieve module:", err)
	}
	socketName := agent.LaunchModule(module)

	sc, err := ipc.StartServer(socketName, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Start ipc server for socket:", socketName)

	agent.WaitConnection(sc, &db)

	//send instruction
	// agent.CheckSendMessage(sc, "world")
	// time.Sleep(3 * time.Second)
	// agent.CheckSendMessage(sc, "toto")
	// time.Sleep(3 * time.Second)
	// resp := agent.CheckSendMessageAndWaitResponse(sc, "youhou")
	// fmt.Println(resp)
	sayHelloInstr := agent.Instruction{Type: agent.Run}
	if err := agent.SendInstruction(sc, sayHelloInstr); err != nil {
		fmt.Println("failed sending instruciton:", err)
	}
	if resp, err := agent.WaitResponse(sc); err != nil {
		fmt.Println("Error while waiting for response:", err)
	} else {
		fmt.Println(resp)
	}
	time.Sleep(2 * time.Second)
	//kill
	agent.KillModule(sc)

}
