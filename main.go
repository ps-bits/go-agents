package main

import (
	"fmt"
	"time"

	"go-agents/agents"
)

type Message struct {
	subject string
	content string
}

type Data struct {
	counter int
}

func (d *Data) Reset() {
	d.counter = 0
}

// Example reader:
func reader(message Message, data *Data) int {
	// Return values:
	// 1	success, do nothing, continue
	// -1	end - stop the agent
	// 0	will be automatically returned on panic

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nPanic. Recovered. Error:\n", r)
			// Perform some action on panic. In this case, reset the data:
			data.Reset()
		}
	}()

	switch message.subject {
	case "crash":
		panic("testing panic")
	case "stop":
		fmt.Println("\nStopping agent")
		return -1
	}

	data.counter++

	fmt.Println("I got a message:", message, data.counter)

	return 1
}

func example() {
	fmt.Println("\nProgram start ----------------")

	manager := agents.NewManager[Message, Data]()

	agentId := manager.NewAgent(reader, &Data{counter: 0}, 8)

	time.Sleep(1 * time.Millisecond)

	manager.SendTo(agentId, Message{subject: "hello", content: "abc"})
	manager.SendTo(agentId, Message{subject: "hello", content: "abc"})
	manager.SendTo(agentId, Message{subject: "crash", content: "abc"})
	time.Sleep(1 * time.Millisecond)
	manager.SendTo(agentId, Message{subject: "hello", content: "abc"})
	manager.SendTo(agentId, Message{subject: "stop", content: "abc"})

	time.Sleep(1 * time.Second)

	fmt.Println("\nProgram end ----------------")
}

func main() {
	example()
}
