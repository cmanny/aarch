package exe

import (
	"github.com/cmanny/aarch/architecture/comp"
	"github.com/cmanny/aarch/architecture/ins"
)

type AuIn struct {
	code byte
	op1  int
	op2  int
}

type ArithmeticUnit struct {
	comp.Communicator
}

func (au *ArithmeticUnit) Init() {
	au.InitComms()
}

func (au *ArithmeticUnit) Data() interface{} {
	return ""
}

func (au *ArithmeticUnit) State() string {
	return ""
}

func (au *ArithmeticUnit) Cycle() {
	op := (<-au.Inputs["in"]).(AuIn)
	res := 0
	switch {
	case op.code == ins.MUL || op.code == ins.MULI:
		res = op.op1 * op.op2
	case op.code == ins.ADD || op.code == ins.ADDI:
		res = op.op1 + op.op2
	case op.code == ins.SUB || op.code == ins.SUBI:
		res = op.op1 - op.op2
	case op.code == ins.XOR || op.code == ins.XORI:
		res = op.op1 ^ op.op2
	}

	au.Outputs["out"] <- res

}
