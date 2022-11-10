package mylogger2

import (
	"log"
)

func LogInfo(message string) {
	log.Printf("Info - %v", message)
}

func LogWarning(message string) {
	log.Printf("Warn - %v", message)
}

func LogError(message string) {
	log.Printf("Error - %v", message)
}
