package despedida

import (
	"fmt"

	"github.com/programatta/midemolib01/internal"
)

// DespedidaSimple returns a single message
func DespedidaSimple() string {
	return "Adios y gracias por usarme"
}

// DespedidaType1 retuns a single message with MsgType1
func DespedidaType1() string {
	m := internal.MsgType1()
	return fmt.Sprintf("Adios y gracias con un msgType1:%s", m)
}
