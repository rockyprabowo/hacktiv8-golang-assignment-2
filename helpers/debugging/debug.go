package debugging

import "log"

// Print prints a debugging message to the console.
func Print(message any, source string) {
	log.Printf("debug: [%s] %s", source, message)
}
