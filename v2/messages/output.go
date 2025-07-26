package messages

type Output interface {
	OutputMessage(stream chan string)
}
