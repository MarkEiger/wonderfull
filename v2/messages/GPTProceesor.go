package messages

import (
	"context"
	"go/v2/tools"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/ssestream"
	"github.com/openai/openai-go/responses"
)

type GPTProcessor struct {
	ApiKey           string
	Model            string
	LastId           string
	Preserve_context bool
}

func BuildParams(processor *GPTProcessor, message string) responses.ResponseNewParams {
	/*
		builds the basic format of the params OpenAi's API requires,
		optionally adds previous response ID for the converstation to have context
	*/
	var toolsDeclarations []responses.ToolUnionParam = []responses.ToolUnionParam{}
	for tool := range tools.Functions {
		var param responses.ToolUnionParam = tools.Functions[tool].GetDeclaration()
		toolsDeclarations = append(toolsDeclarations, param)
	}

	params := responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(message),
		},
		Model: openai.ChatModelGPT4oMini,
		Tools: toolsDeclarations,
	}

	if processor.LastId != "" && processor.Preserve_context {
		params.PreviousResponseID = openai.String(processor.LastId)
	}

	return params
}

func ParseMessage(MessageStream *ssestream.Stream[responses.ResponseStreamEventUnion], outputChannel chan string) {
	for MessageStream.Next() {
		chunk := MessageStream.Current()
		outputChannel <- chunk.Delta.OfString
	}
}

func ParseFunctionCall(FunctionStream *ssestream.Stream[responses.ResponseStreamEventUnion], outputChannel chan string) {
	/*
		parses and executes the desired function call
	*/
	var argumentsBuilder strings.Builder
	var functionName string = FunctionStream.Current().Item.Name
	for FunctionStream.Next() {
		chunk := FunctionStream.Current()
		argumentsBuilder.WriteString(chunk.Delta.OfString)
	}
	tool, ok := tools.Functions[functionName]
	if ok {
		args := argumentsBuilder.String()
		outputChannel <- tool.Perform([]byte(args))
	} else {
		outputChannel <- "Unsupported function: " + functionName
	}
}

func ParseStream(openAiStream *ssestream.Stream[responses.ResponseStreamEventUnion], outputChannel chan string, processor *GPTProcessor) {
	for openAiStream.Next() {
		chunk := openAiStream.Current()
		if chunk.Type == "response.output_item.added" {
			switch chunk.Item.Type {
			case "function_call":
				ParseFunctionCall(openAiStream, outputChannel)
			case "message":
				ParseMessage(openAiStream, outputChannel)
			}
		} else if processor.Preserve_context && chunk.Response.ID != "" {
			processor.LastId = chunk.Response.ID
		}
	}

	if openAiStream.Err() != nil {
		println(openAiStream.Err().Error())
	}
}

func (processor *GPTProcessor) Process(message string, outputChannel chan string) {
	defer close(outputChannel)

	client := openai.NewClient(option.WithAPIKey(processor.ApiKey))
	params := BuildParams(processor, message)

	openAiStream := client.Responses.NewStreaming(
		context.Background(),
		params,
	)

	defer openAiStream.Close()
	ParseStream(openAiStream, outputChannel, processor)
}
