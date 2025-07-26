package tools

import "github.com/openai/openai-go/responses"

type Tool interface {
	Perform(args []byte) string
	GetDeclaration() responses.ToolUnionParam
}
