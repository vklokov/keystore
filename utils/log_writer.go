package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogWriter struct {
	Location string
}

func (l *LogWriter) Info(message string) {
	now := time.Now().UTC()
	l.Write([]byte(fmt.Sprintf("[%v] Info: %v\n", now.Format(DATE_TIME_FORMAT), message)))
}

func (l *LogWriter) Error(err error) {
	now := time.Now().UTC()
	l.Write([]byte(fmt.Sprintf("\n[%v] *Error: %v\n", now.Format(DATE_TIME_FORMAT), err)))
}

func (l *LogWriter) Write(data []byte) (n int, res error) {
	file, err := os.OpenFile(l.Location, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		log.Fatalf("* Error opening file %v, %v", l.Location, err)
	}

	defer file.Close()

	return file.Write(data)
}
