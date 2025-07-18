package l

import (
	"log"
)

var (
	Error *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

func init() {
	Error = log.Default()
	Info = log.Default()
}
