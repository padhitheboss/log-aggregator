package logger

import (
	"fmt"
	"sync"

	loglevel "example.com/log-aggregate/pkg/logLevel"
	"example.com/log-aggregate/pkg/model/message"
	"example.com/log-aggregate/pkg/model/sink"
)

type Logger struct {
	name         string
	sinks        []sink.Sink
	messageQueue chan message.Message
	waitGroup    sync.WaitGroup // Added WaitGroup for synchronization
}

func NewLogger(name string, bufferSize int) *Logger {
	logger := Logger{
		name:         name,
		messageQueue: make(chan message.Message, bufferSize),
	}
	logger.waitGroup.Add(1) // Increment WaitGroup counter
	go func() {
		for {
		}
	}()
	go logger.ForwardMessages()
	return &logger
}

func (l *Logger) PushMessage(log string) {
	message := message.CreateMessage(log)
	l.messageQueue <- message
}

func (l *Logger) AddSink(s sink.Sink) {
	l.sinks = append(l.sinks, s)
}

func (l *Logger) ForwardMessages() {
	// defer l.waitGroup.Done() // Decrement WaitGroup counter when done

	for {
		select {
		case msg := <-l.messageQueue:
			for _, s := range l.sinks {
				slevel := loglevel.LogLevel[s.GetLogLevel()]
				mlevel, isValid := loglevel.LogLevel[msg.GetLogLevel()]
				if !isValid {
					fmt.Println("invalid message")
				}
				if slevel <= mlevel {
					s.PushMessage(msg)
				}
			}
		}
	}
}

func (l *Logger) Wait() {
	l.waitGroup.Wait() // Wait until the WaitGroup counter goes to zero
}
