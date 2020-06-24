package internal

import "fmt"

// MsgType1 creates a single message
func MsgType1() string {
	return "Mensaje tipo 1"
}

// MsgType2 create a message with saludo
func MsgType2(saludo string) string {
	return fmt.Sprintf("Mensaje tipo 2 con saludo: %s", saludo)
}
