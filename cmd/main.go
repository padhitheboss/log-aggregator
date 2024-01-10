package main

import (
	"fmt"
	"time"

	"example.com/log-aggregate/pkg/model/logger"
	"example.com/log-aggregate/pkg/model/sink"
)

func main() {
	logger := logger.NewLogger("logger", 1)
	defer logger.Wait()
	s1, _ := sink.CreateSink("s1", "2006-01-02 15:04:05", "DEBUG", "CONSOLE")
	s2, _ := sink.CreateSink("s2", "2006-01-02 15:04:05", "ERROR", "CONSOLE")
	s3, _ := sink.CreateSink("s3", "2006-01-02 15:04:05", "WARN", "CONSOLE")
	s4, err := sink.CreateSink("s4", "2006-01-02 15:04:05", "INFO", "CONSOLE")
	if err == nil {
		logger.AddSink(s1)
		logger.AddSink(s2)
		logger.AddSink(s3)
		logger.AddSink(s4)
	} else {
		fmt.Println(err)
	}
	logger.PushMessage("INFO Application is Running on Port 8080")
	time.Sleep(1 * time.Second)
	logger.PushMessage("DEBUG User Connection established")
	time.Sleep(2 * time.Second)
	logger.PushMessage("WARN Possible Memory leak detected")
	time.Sleep(4 * time.Second)
	logger.PushMessage("ERROR Worker 1 is not responding")
}
