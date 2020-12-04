package netutils

import (
	"errors"
	"log"
	"strconv"
)

func NormalizePort(value string) ConnectionTarget {
	ret := ConnectionTarget{}
	port, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		if errors.Is(err, strconv.ErrRange) {
			log.Fatalln("port out of range 0-65535", value)
		}
		if errors.Is(err, strconv.ErrSyntax) { // named pipe
			ret.TargetType = NamedPipe
			ret.NamedPipe = value
			return ret
		}
	}
	ret.TargetType = Port
	ret.Port = uint16(port)
	return ret
}
