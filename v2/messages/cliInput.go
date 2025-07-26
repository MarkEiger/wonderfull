package messages

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLIInput struct {
}

func (Input *CLIInput) GetMesage() string {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error encountered: ", err)
		return ""
	}
	return strings.TrimSpace(input)
}
