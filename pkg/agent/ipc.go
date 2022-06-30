package agent

import (
	"fmt"

	ipc "github.com/james-barrow/golang-ipc"
)

//CheckSendMessage: send message and check errors
func CheckSendMessage(sc *ipc.Server, msg string) {
	if err := sc.Write(5, []byte(msg)); err != nil {
		fmt.Println("Error while sending message:", err)
	}
}

//CheckSendMessageAndWaitResponse:send message wait for response and check errors
func CheckSendMessageAndWaitResponse(sc *ipc.Server, msg string) (response string) {
	if err := sc.Write(5, []byte(msg)); err != nil {
		fmt.Println("Error while sending message:", err)
	}
	response, err := WaitResponse(sc)
	if err != nil {
		fmt.Println("Error while waiting IPC response:", err)
		return response
	}

	return response
}

func WaitResponse(sc *ipc.Server) (response string, err error) {
	for {
		resp, err := sc.Read()
		if err == nil {
			response = string(resp.Data)
			return response, nil
		} else {
			return response, err
		}
	}
}
