package message

import (
	"strings"
	"time"
)

type Message struct {
	message   string
	timeStamp time.Time
	level     string
}

func CreateMessage(message string) Message {
	msgAttrib := strings.Split(message, " ")
	lvl := msgAttrib[0]
	msg := strings.Join(msgAttrib[1:], " ")
	return Message{message: msg, timeStamp: time.Now(), level: lvl}
}

func (m *Message) String(timeFormat string) string {
	return m.timeStamp.Format(timeFormat) + " " + m.level + " " + m.message
}

func (m *Message) GetLogLevel() string {
	return m.level
}
