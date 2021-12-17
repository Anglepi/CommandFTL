package logger

import (
	"log"
	"os"
)

var (
	MyLog *log.Logger
	file  *os.File
	err   error
)

func InitLog() {
	file, err = os.OpenFile("FTL.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	MyLog = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func CloseLog() {
	file.Close()
}
