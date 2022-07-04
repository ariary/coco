package main

import (
	"github.com/ariary/coco/pkg/server"
	prompt "github.com/c-bata/go-prompt"
)

func main() {
	defer server.HandleExit()

	ctx := &server.CocoContext{}

	p := prompt.New(
		server.Executor(ctx),
		server.Completer,
		prompt.OptionLivePrefix(server.LivePrefix),
		prompt.OptionTitle("coco"),
	)
	p.Run()
}
