package optimizer

import (
	"log"
	"os"
)

var DebugLogger *log.Logger

func InitLoggers() {
	DebugLogger = log.New(os.Stdout, "", log.Lshortfile)
}