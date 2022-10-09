package debugging

import "log"

// Print prints a debugging message to the console.
func Print(message string, source string) {
	log.Printf("debug: [%s] %s", source, message)
}
