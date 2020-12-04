package main

import (
	"bufio"
	"fmt"
	"gopkg.in/natefinch/npipe.v2"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"utils/src/info"
	"utils/src/netutils"
)

const receiveBufferSizeKey = "LISTENER_RECEIVE_BUFFER_SIZE"
const defaultReceiveBufferSize = 32

func main() {
	info.CheckPrintInfo()
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("got %v arguments, expected 1\n", len(args))
	}
	var err error
	var receiveBufferSize = defaultReceiveBufferSize
	size := os.Getenv(receiveBufferSizeKey)
	if size != "" {
		receiveBufferSize, err = strconv.Atoi(size)
		if err != nil {
			log.Printf("unable to parse %v, using default value (%v)", receiveBufferSizeKey, defaultReceiveBufferSize)
			receiveBufferSize = defaultReceiveBufferSize
		}
	}

	target := netutils.NormalizePort(args[0])

	var connection net.Conn

	switch target.TargetType {
	case netutils.Port:
		connection, err = net.Dial("tcp", fmt.Sprintf(":%v", target.Port))
		if err != nil {
			log.Fatalln("unable to open port", err)
		}
	case netutils.NamedPipe:
		connection, err = npipe.Dial(target.NamedPipe)
		if err != nil {
			log.Fatalln("unable to open named pipe", err)
		}
	default:
	case netutils.Undefined:
		log.Fatalln("no listen target provided")
	}
	defer func() {
		err := connection.Close()
		if err != nil {
			log.Fatalln("unable to close connection", err)
		}
	}()
	buffer := make([]byte, receiveBufferSize)
	reader := bufio.NewReader(connection)
	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				return // exit
			} else {
				log.Fatalln(err)
			}
		}
		if bytesRead > 0 {
			fmt.Print(string(buffer[:bytesRead]))
		}
	}
}
