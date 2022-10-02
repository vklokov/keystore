package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vklokov/keystore/utils"
)

var Logger *LogWriter

func CreateLogger(location string) *LogWriter {
	Logger = &LogWriter{
		Location: location,
	}

	return Logger
}

type LogWriter struct {
	Location string
}

func (self *LogWriter) Write(data []byte) (n int, res error) {
	file, err := os.OpenFile(self.Location, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		log.Fatalf("* Error opening file %v, %v", self.Location, err)
	}

	defer file.Close()
	return file.Write(data)
}

func (l *LogWriter) Info(message string) {
	now := time.Now().UTC()
	l.Write([]byte(fmt.Sprintf("[%v] Info: %v\n", now.Format(utils.DATE_TIME_FORMAT), message)))
}

func (l *LogWriter) Error(err error) {
	now := time.Now().UTC()
	l.Write([]byte(fmt.Sprintf("\n[%v] *Error: %v\n", now.Format(utils.DATE_TIME_FORMAT), err)))
}
