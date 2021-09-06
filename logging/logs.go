package logging

import (
	"log"
	"os"
)

var Info, Warning, Error *log.Logger

func init() {
	Info = log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "[WARNING]", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "[ERRROR]", log.Ldate|log.Ltime|log.Lshortfile)
}