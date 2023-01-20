package agents

import "fmt"

func NewManager[T any, D any]() Manager[T, D] {
	m := Manager[T, D]{
		list:    make(map[int]chan T),
		counter: 0,
	}
	return m
}

type Manager[T any, D any] struct {
	list    map[int]chan T
	counter int
}

func (m *Manager[T, D]) NewAgent(reader func(T, *D) int, data *D, chanSize int) int {
	// Returns the id of the created agent.
	ch := make(chan T, chanSize)
	m.list[m.counter] = ch
	m.counter++
	go messageLoop[T](ch, reader, data)
	return m.counter - 1
}

func (m *Manager[T, D]) SendTo(id int, msg T) bool {
	// Tries to send a message to agent with id.
	select {
	case m.list[id] <- msg:
		return true
	default:
		return false
	}
}

func messageLoop[T any, D any](main chan T, reader func(T, D) int, data D) {
	// This will be run as a go-routine.
	for {
		select {
		case mes := <-main:
			result := reader(mes, data)
			switch result {
			case 1:
				// Ok
			case 0:
				fmt.Println("Detected panic. No action.")
			case -1:
				// End this go-routine.
				return
			}
		}
	}
}
