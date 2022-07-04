package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

//// All message types (TextMessage, BinaryMessage, CloseMessage, PingMessage and
// PongMessage) are supported.

func read(end chan struct{}, c *websocket.Conn) {
	defer close(end)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}
func main() {
	var addr = flag.String("c", "localhost:9292", "Coco C2 server address")
	flag.Parse()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/agent"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	kill := make(chan struct{})

	go read(kill, c)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-kill:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-kill:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/ariary/coco/pkg/agent"
// 	ipc "github.com/james-barrow/golang-ipc"
// )

// func main() {
// 	var url string
// 	flag.StringVar(&url, "u", "", "url testing purpose")
// 	flag.Parse()

// 	db := agent.ModuleDB{}

// 	module, err := agent.GetModuleContentHTTP(url)
// 	if err != nil {
// 		fmt.Println("failed to retrieve module:", err)
// 	}

// 	socket := agent.LaunchModule(module)

// 	sc, err := ipc.StartServer(socket, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println("Start ipc server for socket:", socket)

// 	agent.WaitConnection(sc, &db)

// 	//send instruction
// 	// agent.CheckSendMessage(sc, "world")
// 	// time.Sleep(3 * time.Second)
// 	// agent.CheckSendMessage(sc, "toto")
// 	// time.Sleep(3 * time.Second)
// 	// resp := agent.CheckSendMessageAndWaitResponse(sc, "youhou")
// 	// fmt.Println(resp)
// 	sayHelloInstr := agent.Instruction{Type: agent.Run}
// 	if err := agent.SendInstruction(sc, sayHelloInstr); err != nil {
// 		fmt.Println("failed sending instruciton:", err)
// 	}
// 	if resp, err := agent.WaitResponse(sc); err != nil {
// 		fmt.Println("Error while waiting for response:", err)
// 	} else {
// 		fmt.Println(resp)
// 	}
// 	time.Sleep(2 * time.Second)
// 	//kill
// 	agent.KillModule(sc)

// }

// //TODO: pass parameter + websocket server
