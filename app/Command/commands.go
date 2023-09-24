package Commands

import "strings"

type RedisCommand byte

const (
	GET RedisCommand = iota
	SET
	PING
	COMMAND
	DOCS
	UNKNOW
)

type Command struct {
	Cmd   RedisCommand
	Key   *string
	Value *string
	Ttl   *int
}

func NewCommand(cmd string, key string, value string, ttl int) Command {

	switch strings.ToLower(cmd) {
	case "ping":
		return Command{PING, nil, nil, nil}
	case "command":
		return Command{COMMAND, nil, nil, nil}
	case "docs":
		return Command{DOCS, nil, nil, nil}
	case "get":
		return Command{GET, &key, nil, nil}
	case "set":
		return Command{SET, &key, &value, &ttl}
	default:
		return Command{UNKNOW, nil, nil, nil}
	}
}

func (cmd *Command) String() string {
	switch cmd.Cmd {
	case PING:
		return "PING"
	case GET:
		return "GET"
	case SET:
		return "SET"
	case COMMAND:
		return "COMMAND"
	case DOCS:
		return "DOCS"
	default:
		return "UNKNOW"
	}
}
