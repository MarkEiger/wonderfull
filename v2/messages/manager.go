package messages

import (
	"sync"
)

type Manager struct {
	Input     Input
	Processor Processor
	Output    Output
	WaitGroup *sync.WaitGroup
}

func (manager *Manager) Manage() {
	for {
		message := manager.Input.GetMesage()
		if message == "exit" {
			return
		}
		outputChannel := make(chan string)
		go manager.Processor.Process(message, outputChannel)
		manager.WaitGroup.Add(1)
		go manager.Output.OutputMessage(outputChannel)
		manager.WaitGroup.Wait()
	}
}
