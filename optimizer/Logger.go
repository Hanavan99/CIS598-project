package optimizer

import (
	"log"
	"os"
	"io/ioutil"
)

var DebugLogger *log.Logger

func InitLoggers() {
	DebugLogger = log.New(os.Stdout, "", log.Lshortfile)
	DebugLogger.SetOutput(ioutil.Discard)
}