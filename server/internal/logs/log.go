package logs

import (
	"log"
	"os"
)

var LogErrorLogger = log.New(os.Stdout, "[ADON][ERROR]", log.LstdFlags)
var LogInfoLogger = log.New(os.Stdout, "[ADON][INFO]", log.LstdFlags)
