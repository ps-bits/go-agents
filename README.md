# go-agents
A package that gives an alternative flavor to Go's concurrency.  
It's isnpired by the 'actor model' found in some programming languages.

## Usage
Agents are go-routines that hold DATA and processes MESSAGES using a READER function.  
Before creating an Agent, you first need to declare a Data type, a Message type, and a reader.  
You also need a "manager". It keeps track of multiple agents, and gives each a unique id.

```go
// Declare a message type - agents communicate with the outside world with messages
type Message struct {
	body string
}

// Declare a Data type - the agents internal state
type Data struct {
	counter int
}

// Declare a reader - defines how the agents processes messages and modifies its data
func reader(message Message, data *Data) int {
	data.counter++
	fmt.Println("I this message:", message.body, " for the ", data.counter, "th time!")
	return 1 // return 1 for success, return -1 to stop the Agent process
}

// A manager holds a list of agents, and keeps track of their id's
manager := agents.NewManager[Message, Data]()

// The new agent needs a reader, an initial data, and a size for its inbox
// The agent will have a unique id. Used when sending messages to it.
agentId := manager.NewAgent(reader, &Data{counter: 0}, 8)

// Sends a message to the agent
manager.SendTo(agentId, Message{body: "Hello!"})

```
