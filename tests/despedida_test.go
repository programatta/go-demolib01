package tests

import (
	"testing"

	"github.com/programatta/midemolib01/pkg/despedida"
)

func TestDespedidaSimple(t *testing.T) {
	expected := "Adios y gracias por usarme"
	if msg := despedida.DespedidaSimple(); msg != expected {
		t.Errorf("DespedidaSimple() returns \"%s\" and expected \"%s\"", msg, expected)
	}
}
