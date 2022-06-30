package c2c

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	ipc "github.com/james-barrow/golang-ipc"
)

type ModuleDB struct {
	Modules []Module
}

type Module struct {
	Socket *ipc.Server
	Name   string
}

type InstructionType int64

const (
	Run InstructionType = iota
	Stop
	Kill
)

type Params struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Instruction struct {
	Type   InstructionType `json:"type"`
	Params []Params        `json:"paremeters"`
}

//WaitConnection: Wait for module connexion
func WaitConnection(sc *ipc.Server, db *ModuleDB) {
	//wait connection
	for {
		msg, err := sc.Read()
		if err == nil {
			msgStr := string(msg.Data)
			if strings.HasPrefix(msgStr, CONNECTION_KEYWORD) {
				moduleName := strings.Split(msgStr, ":")[len(strings.Split(msgStr, ":"))-1] //avoid testing size and reassignement -_-
				fmt.Println("🛰️ module", moduleName, "loaded")
				//confirm connection
				CheckSendMessage(sc, LOADED_KEYWORD)
				module := Module{Name: moduleName, Socket: sc}
				db.Modules = append(db.Modules, module)
				break
			}
		} else {
			log.Println(err)
			// return
		}

	}
}

func SendInstruction(sc *ipc.Server, instr Instruction) (err error) {
	//Send instruction
	instrJSON, err := json.Marshal(instr)
	if err != nil {
		return err
	}
	if err := sc.Write(5, instrJSON); err != nil {
		return err
	}

	//ACK
	if resp, err := WaitResponse(sc); err != nil {
		return err
	} else if strings.Contains(resp, INSTR_OK) {
		return errors.New("do not received instruction ACK: " + resp)
	}
	return nil
}

//  response := money{}
// json.Unmarshal([]byte(answer), &response)
