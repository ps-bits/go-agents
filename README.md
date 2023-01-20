# go-agents
Go concurrency package.
  
This is to make it easier to use concurrency.
It is an alternative flavor to go's original way.

### Agent
An agent represents a concurrent process. I has associated Data, Message, and Message reader.
Agents read Messages sent to them one by one. They have an internal state called Data.

### Data
The data for the agent. Can be whatever you define.   
Ideally it should be self-contained: it should not have pointers pointin to or from it.

### Message
The messages that the agent should handle.

### Message Reader
A function that reads and handles the messages given to the agent.

### Manager
Has a list of Agents.
The Agents can't be accessed directly. Only via an "id".
The manager makes sure that every Agent has a unique id.
Id's can not be reused.



### Note on Panicked functions
A panicked function will normally return 0/nil/false.  
(But can return other values when named return params are used)

https://stackoverflow.com/questions/33167282/how-to-return-a-value-in-a-go-function-that-panics