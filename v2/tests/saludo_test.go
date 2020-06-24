package tests

import (
	"fmt"
	"testing"

	"github.com/programatta/midemolib01/internal"
	saludov2 "github.com/programatta/midemolib01/v2/pkg/saludo"
)

func TestMiSaludoSimple(t *testing.T) {
	expected := "Hola simple"
	if msg := saludov2.MiSaludoSimple(); msg != expected {
		t.Errorf("MiSaludoSimple() returns %s and expected %s", msg, expected)
	}
}

func TestMiSaludoType1(t *testing.T) {
	msg1 := internal.MsgType1()
	expected := fmt.Sprintf("Mi saludo con un msgType1:%s", msg1)
	if msg := saludov2.MiSaludoType1(); msg != expected {
		t.Errorf("MiSaludoType1() returns %s and expected %s", msg, expected)
	}
}

func TestMiSaludoType2(t *testing.T) {
	msg := "Hola caracola"
	msg2 := internal.MsgType2(msg)
	expected := fmt.Sprintf("Mi saludo con un msgType2:%s", msg2)
	if msgret := saludov2.MiSaludoType2(msg); msgret != expected {
		t.Errorf("MiSaludoType2('%s') returns %s and expected %s", msg, msgret, expected)
	}
}
