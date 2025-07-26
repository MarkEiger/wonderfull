package arguments

import (
	"fmt"
	"os"
)

func ParseArguments(args []string) string {
	if len(args) == 0 {
		fmt.Println("No Arguments supplied, aborting")
		return ""
	}
	for index, arg := range args {
		switch arg {
		case "-h":
			fmt.Println(`
-h     get this help message
-k     OpenAI API Key
-v     Use Enviromental Variale OPENAI_API_KEY for the key
to exit the chat enter the word 'exit'
			`)
			return ""
		case "-k":
			if len(args) < index+2 {
				fmt.Println("No Key Supplied, aborting")
				return ""
			}
			return args[index+1]
		case "-v":
			key := os.Getenv("OPENAI_API_KEY")
			if key == "" {
				fmt.Println("Enviromental variable OPENAI_API_KEY not defined")
			}
			return key
		default:
			fmt.Println("Unsupported Argument " + arg + " , Use -h for help")
			return ""
		}
	}
	return ""
}
