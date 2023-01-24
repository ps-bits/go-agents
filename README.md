# go-agents
A package that gives an alternative flavor to Go's concurrency.

## Usage
Agents are go-routines that hold data (an internal state), and processes messages, using a function called a "reader".  
To create an Agent, you first need to declare a Message type, a Data type, and a reader.  
Then create a manager, and then finally the Agent.  

```go
type Message struct {
	subject string
	content string
}

type Data struct {
	counter int
}

func reader(message Message, data *Data) int {
	data.counter++
	fmt.Println("I got a message:", message, " for the ", data.counter, "th time!")
	return 1
}

manager := agents.NewManager[Message, Data]()

agentId := manager.NewAgent(reader, &Data{counter: 0}, 8)

```