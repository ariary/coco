package main

import (
	"fmt"
	"log"
	"strings"

	ipc "github.com/james-barrow/golang-ipc"
)

const CONNECTION_KEYWORD = "CONNECTED"
const LOADED_KEYWORD = "LOADED"
const SOCKET = "YAMANAKA"

func main() {
	sc, err := ipc.StartServer(SOCKET, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Start ipc server for socket:", SOCKET)

	//wait connection
	for {
		msg, err := sc.Read()
		if err == nil {
			// fmt.Printf("%+v", data)
			msgStr := string(msg.Data)
			if strings.HasPrefix(msgStr, CONNECTION_KEYWORD) {
				module := strings.Split(msgStr, ":")[len(strings.Split(msgStr, ":"))-1] //avoid testing size and reassignement -_-
				fmt.Println("🛰️ module", module, "loaded")
				//confirm connection
				CheckSendMessage(sc, LOADED_KEYWORD)

				//add socket to a Struct
				break
			}
		} else {
			log.Println(err)
			return
		}

	}

	//send instruction
	CheckSendMessage(sc, "world")
	//TODO: for loop to wait response

}

//goroutin wait response?

//CheckSendMessage: send message and check errors
func CheckSendMessage(sc *ipc.Server, msg string) {
	if err := sc.Write(5, []byte(msg)); err != nil {
		fmt.Println("Error while sending message:", err)
	}
}
