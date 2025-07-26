package tools

import (
	"encoding/json"
	"strconv"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/responses"
)

type MultiplyToolArgs struct {
	X int `json:"x"`
	Y int `json:"y"`
}

/*
OpenAI API to define a function for the possibilty of a function call
*/
var MultiplyToolDeclaration responses.ToolUnionParam = responses.ToolUnionParam{
	OfFunction: &responses.FunctionToolParam{
		Name:        "Multiply",
		Description: openai.String("The Multiply function multiplies two numbers by one another and returns the result"),
		Parameters: openai.FunctionParameters{
			"type": "object",
			"properties": map[string]interface{}{
				"x": map[string]string{
					"type":        "number",
					"description": "first number to multiply",
				},
				"y": map[string]string{
					"type":        "number",
					"description": "second number to multiply",
				},
			},
			"required": []string{"x", "y"},
		},
	},
}

func Multiply(x int, y int) int {
	return x * y
}

type MultiplyTool struct {
}

func (tool *MultiplyTool) GetDeclaration() responses.ToolUnionParam {
	return MultiplyToolDeclaration
}

func (tool *MultiplyTool) Perform(args []byte) string {
	var getArgs MultiplyToolArgs
	json.Unmarshal([]byte(args), &getArgs)
	return "The resault is: " + strconv.Itoa(getArgs.X*getArgs.Y)
}
