package main

import (
	"go/v2/arguments"
	"go/v2/messages"
	"os"
	"sync"
)

func main() {
	key := arguments.ParseArguments(os.Args[1:])
	if key == "" {
		return
	}
	waitGroup := sync.WaitGroup{}
	reader := &messages.CLIInput{}
	processor := &messages.GPTProcessor{
		ApiKey:           key,
		Model:            "gpt-4o-mini",
		LastId:           "",
		Preserve_context: false,
	}
	output := &messages.StreamOutput{WaitGroup: &waitGroup}
	m := messages.Manager{
		Input:     reader,
		Output:    output,
		Processor: processor,
		WaitGroup: &waitGroup,
	}
	m.Manage()
}
