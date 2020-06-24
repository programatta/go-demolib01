package tests

import (
	"testing"

	despedidav2 "github.com/programatta/midemolib01/v2/pkg/despedida"
)

func TestDespedidaSimple(t *testing.T) {
	expected := "Adios y gracias por usarme"
	if msg := despedidav2.DespedidaSimple(); msg != expected {
		t.Errorf("DespedidaSimple() returns \"%s\" and expected \"%s\"", msg, expected)
	}
}
