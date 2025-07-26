package messages

import (
	"fmt"
	"sync"
)

type StreamOutput struct {
	WaitGroup *sync.WaitGroup
}

func (stream *StreamOutput) OutputMessage(outputStream chan string) {
	for message := range outputStream {
		fmt.Print(message)
	}
	fmt.Println()
	stream.WaitGroup.Done()
}
