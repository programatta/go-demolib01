package tests

import (
	"testing"

	despedidav2 "github.com/programatta/midemolib01/pkg/despedida/v2"
)

func TestDespedidaSimple(t *testing.T) {
	expected := "Adios y gracias por usarme"
	if msg := despedidav2.DespedidaSimple(); msg != expected {
		t.Errorf("DespedidaSimple() returns \"%s\" and expected \"%s\"", msg, expected)
	}
}
