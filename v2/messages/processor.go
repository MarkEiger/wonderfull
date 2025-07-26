package messages

type Processor interface {
	Process(message string, outputChannel chan string)
}
