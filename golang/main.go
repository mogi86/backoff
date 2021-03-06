package main

import "fmt"

type Message string

func InitMessage() Message {
	return Message("hello!")
}

type Logger struct {
	Message Message
}

func InitLogger(m Message) Logger {
	return Logger{Message: m}
}

func main() {
	l := InitLoggerViaWire()
	fmt.Println(l.Message)
}
