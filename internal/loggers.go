package internal

import (
	"log"
	"os"
)

var InfoLog = log.New(os.Stdout, "INFO\t", log.LUTC|log.Ldate|log.Ltime)

var ErrorLog = log.New(os.Stderr, "ERROR\t", log.LUTC|log.Ldate|log.Ltime|log.Llongfile)
