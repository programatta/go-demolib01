package tests

import (
	"fmt"
	"testing"

	"github.com/programatta/midemolib01/internal"
)

func TestMsgType1(t *testing.T) {
	expedted := "Mensaje tipo 1"
	if msg := internal.MsgType1(); msg != expedted {
		t.Errorf("MsgType1() returns %s and expected %s", msg, expedted)
	}
}

func TestMsgType2(t *testing.T) {
	saludo := "Hola hola caracola"
	expedted := fmt.Sprintf("Mensaje tipo 2 con saludo: %s", saludo)
	if msg := internal.MsgType2(saludo); msg != expedted {
		t.Errorf("MsgType2() returns %s and expected %s", msg, expedted)
	}
}
