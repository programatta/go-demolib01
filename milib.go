package midemolib01

import (
	"fmt"

	"github.com/programatta/midemolib01/internal"
)

// MiSaludoSimple retunrs a single greeting
func MiSaludoSimple() string {
	return "Hola simple"
}

// MiSaludoType1 retuns a single greeting with MsgType1
func MiSaludoType1() string {
	m := internal.MsgType1()
	return fmt.Sprintf("Mi saludo con un msgType1:%s", m)
}

// MiSaludoType2 returns a greeting with msg using MsgType2
func MiSaludoType2(msg string) string {
	m := internal.MsgType2(msg)
	return fmt.Sprintf("Mi saludo con un msgType2:%s", m)
}
