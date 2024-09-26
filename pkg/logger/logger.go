package logger

import (
	"log"
	"os"
)

type Loggers struct {
	Info  *log.Logger
	Debug *log.Logger
}

func NewLoggers() *Loggers {
	return &Loggers{
		Info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Debug: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
