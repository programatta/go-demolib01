package tests

import (
	"fmt"
	"testing"

	"github.com/programatta/midemolib01/internal"
	despedidav2 "github.com/programatta/midemolib01/v2/pkg/despedida"
)

func TestDespedidaSimple(t *testing.T) {
	expected := "Adios y gracias por usarme"
	if msg := despedidav2.DespedidaSimple(); msg != expected {
		t.Errorf("DespedidaSimple() returns \"%s\" and expected \"%s\"", msg, expected)
	}
}

func TestDespedidaType1(t *testing.T) {
	msg1 := internal.MsgType1()
	expected := fmt.Sprintf("Adios y gracias con un msgType1:%s", msg1)
	if msg := despedidav2.DespedidaType1(); msg != expected {
		t.Errorf("DespedidaSimple() returns \"%s\" and expected \"%s\"", msg, expected)
	}
}
