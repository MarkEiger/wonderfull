# 🚀 Getting Started
To run the CLI tool, follow these steps:

## Navigate to the Project Directory:
```bash
cd v2
```
## Run the Application:
Use the following command to start the application:
```bash
go run main.go [-k key|-v]
```
# 🔑 Passing the API Key
You can pass your API key in two ways:

- Using Command Line Argument:
```bash
-k key
```
- Using Environment Variable:
Set the OPENAI_API_KEY environment variable:
```bash
export OPENAI_API_KEY=your_api_key
```
Then run the application with the -v flag:
```bash
go run main.go -v
```

# ❌ Exiting the Chat
To exit the chat, simply type:

```
exit
```

# 💬 Features
- Question Answering: The chat supports answering a variety of questions.
- Multiplication Function: You can multiply two numbers using the defined function.
# 🛠️ Extending Functionality
To expand the capabilities of the CLI tool, you can add more functions by implementing the interface defined in [v2/tools/tool.go].

    1. Locate the File:
    Open the file located at [v2/tools/functions.go].
    2. Implement New Functions:
    Add your desired functions that adhere to the interface specifications. Ensure that each function is well-documented for clarity.
