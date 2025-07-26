To run the Cli simply:
```
cd v2
go run main.go [-k key|-v]
```

to pass a key use either
- `-k` key
- `-v` with an enviromental variable OPENAI_API_KEY

to exit the chat simply type exit


the chat supports answering questions and a defined function for multiplying two numbers.

to expand on the code add more functions to [v2/tools/functions.go] implementing the interface from [v2/tools/tool.go]
