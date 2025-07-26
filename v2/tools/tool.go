package tools

import "github.com/openai/openai-go/responses"

// the basic struct for a function calling tool
type Tool interface {
	Perform(args []byte) string
	GetDeclaration() responses.ToolUnionParam
}
