# go-agents
A package that gives an alternative flavor to Go's concurrency.

## Usage
Agents are go-routines that hold data (an internal state), and processes messages, using a function called a "reader".  
To create an Agent, you first need to declare a Message type, a Data type, and a reader.  
Then create a manager, and then finally the Agent.  
(The manager gives each agent a unique id)

```go
type Message struct {
	body string
}

type Data struct {
	counter int
}

func reader(message Message, data *Data) int {
	data.counter++
	fmt.Println("I this message:", message.body, " for the ", data.counter, "th time!")
	return 1 // return 1 for success
}

manager := agents.NewManager[Message, Data]()

// The new agent needs a reader, an initial data, and a size for its inbox.
// The agent will have a unique id to identify it. Used when sending messages to it.
agentId := manager.NewAgent(reader, &Data{counter: 0}, 8)

// Sends a message to the agent.
manager.SendTo(agentId, Message{body: "Hello!"})

```

## Reader return values
1: success
-1: stops the process