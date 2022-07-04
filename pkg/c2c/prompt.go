package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

type CocoContext struct {
	C2Servers []CocoServer
	CurrentC2 CocoServer
}

var ctx *CocoContext

// suggestions list
var suggestions = []prompt.Suggest{
	// General
	{"exit", "Exit coco"},
	{"help", "get help method"},

	// key
	{"launch", "launch coco C2 server"},
	// {"keyprint", "print the c serverurrent key"},

	// // Command on ubac
	// {"connect", "Connect to the configured Ubac"}, //in fact launch get and see if there is result
	// {"cd", "Change the  working directory in encrypted fs. (Do not support full path)"},

	// // Read Method
	// {"ls", "list directory contents on remote encrypted fs"},
	// {"cat", "print file content on remote encrypted fs resource"},
	// {"tree", "print tree of remote encrypte fs"},

	// // Write Method
	// {"rm", "remove directory or file on remote encrypted fs"},
}

//Prefix for the prompt
func LivePrefix(ctx CocoContext) (string, bool) {
	return "( coco " + ctx.CurrentC2.Port + ") » ", true
}

//perform at each loop
func Executor(in string, ctx *CocoContext) {
	in = strings.TrimSpace(in)

	// var method, body string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "launch":
		if len(blocks) < 2 {
			fmt.Println("please enter the listenign port of coco C2 server")
		} else {
			ctx.CurrentC2.Port = blocks[1]
			log.SetFlags(0)
			http.HandleFunc("/agent", AgentWebsocket)
			fmt.Println("🥥 Coco C2 server launching on port", ctx.CurrentC2.Port)
			go log.Fatal(http.ListenAndServe(":"+ctx.CurrentC2.Port, nil))
		}
		return
	case "help":
		fmt.Println("available commands: launch")
		return
	case "exit":
		fmt.Println("Bye!🕶")
		HandleExit()
		os.Exit(0)
	default:
		fmt.Printf("Unknown command: %s", blocks[0])
		fmt.Println()
		return
	}

}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

//Function launch when exit. Mainly use to prevent https://github.com/c-bata/go-prompt/issues/228
func HandleExit() {
	rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()
}
