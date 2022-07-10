package logs

import (
	"log"
	"os"
)

var ErrorLogger = log.New(os.Stdout, "[ADON][ERROR]", log.LstdFlags)
var InfoLogger = log.New(os.Stdout, "[ADON][INFO]", log.LstdFlags)
