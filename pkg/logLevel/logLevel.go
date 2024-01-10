package loglevel

var LogLevel map[string]int

func init() {
	LogLevel = make(map[string]int)
	LogLevel["DEBUG"] = 1
	LogLevel["INFO"] = 2
	LogLevel["WARN"] = 3
	LogLevel["ERROR"] = 4
	LogLevel["FATAL"] = 5
}
