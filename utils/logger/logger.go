package logger

import (
	"log"
)

func FatalLog(layer, message string, err error) {
	log.Fatalf("[%s] %s. ERR:%s\n", layer, message, err.Error())
}

func ErrLog(layer, message string, err error) {
	log.Printf("[%s] %s. ERR:%s\n", layer, message, err.Error())
}

func Logger(layer, message string) {
	log.Printf("[%s] %s\n", layer, message)
}