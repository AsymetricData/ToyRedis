package RedisParser

import (
	"bytes"
	Commands "main/app/Command"
)

func ParseBuffer(buffer []byte) []Commands.Command {
	var commands []Commands.Command

	end := bytes.Index(buffer, []byte("\r\n"))

	if end == -1 {
		panic("End line not present")
	}

	switch buffer[0] {
	case '*':
		commands = append(commands, parseArray(buffer[end+2:], int(buffer[1])-48))
	default:
		println("Unknown")
	}

	return commands
}

func parseArray(buffer []byte, len int) Commands.Command {

	commandLength := 0
	valueLength := 0

	command := ""
	key := ""

	buf := buffer

	for index := 0; index <= len+1; index++ {
		lEnd := bytes.Index(buf, []byte("\r\n"))
		lBuf := buf[0:lEnd]

		if lEnd == -1 {
			break
		}

		if commandLength > 0 && valueLength == 0 && command == "" {
			command = string(lBuf[0:commandLength])
		} else if commandLength > 0 && valueLength > 0 {
			key = string(lBuf[0:valueLength])

			return Commands.NewCommand(command, key, "", 0)
		} else {
			switch lBuf[0] {
			case '$':
				//Size of something
				if commandLength == 0 && valueLength == 0 {
					commandLength = int(lBuf[1]) - 48
				} else {
					valueLength = int(lBuf[1]) - 48
				}
			default:
				println("Unhandled lBuf ", string(lBuf[0]))
			}
		}

		buf = buf[lEnd+2:]
	}

	return Commands.Command{}
}
