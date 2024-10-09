package main

import (
	"os"

	"github.com/Harichandra-Prasath/LogIt"
)

var Logger = GetLogger()

func main() {

	server := NewServer(":3000")
	server.SetRoutes()
	err := server.Serve()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *LogIt.Logger {

	opts := LogIt.LoggerOptions{
		Level: LogIt.LEVEL_DEBUG,
		RecordOptions: LogIt.RecordOptions{
			Spacing:   2,
			Colorfull: true,
		},
	}

	logger := LogIt.NewLogger(opts, LogIt.NewTextHandler(os.Stdout, os.Stderr))
	return logger
}
