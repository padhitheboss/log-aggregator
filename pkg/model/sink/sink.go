package sink

import (
	"errors"
	"fmt"

	loglevel "example.com/log-aggregate/pkg/logLevel"
	"example.com/log-aggregate/pkg/model/message"
)

type Sink interface {
	PushMessage(message.Message)
	GetLogLevel() string
}
type ConsoleSink struct {
	name       string
	timeFormat string
	baseLevel  string
}

func CreateSink(name, timeFormat, level, sinkType string) (Sink, error) {
	switch sinkType {
	case "CONSOLE":
		return &ConsoleSink{name, timeFormat, level}, nil
	default:
		return nil, errors.New("invalid sink type: " + sinkType)
	}
}
func (s *ConsoleSink) PushMessage(message message.Message) {
	if loglevel.LogLevel[message.GetLogLevel()] >= loglevel.LogLevel[s.baseLevel] {
		fmt.Println(message.String(s.timeFormat))
	}
}

func (s *ConsoleSink) GetLogLevel() string {
	return s.baseLevel
}
