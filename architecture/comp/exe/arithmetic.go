package exe

import (
	"github.com/cmanny/aarch/architecture/comp"
	"github.com/cmanny/aarch/architecture/ins"
)

type ArithmeticUnit struct {
	comp.Communicator
	currentIns InsIn
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
	res := 0
	for {
		select {
		case _in := <-au.Inputs["in"]:
			in := _in.(InsIn)
			switch {
			case in.Code == ins.MUL || in.Code == ins.MULI:
				res = in.Op2 * in.Op3
			case in.Code == ins.ADD || in.Code == ins.ADDI:
				res = in.Op2 + in.Op3
			case in.Code == ins.SUB || in.Code == ins.SUBI:
				res = in.Op2 - in.Op3
			case in.Code == ins.XOR || in.Code == ins.XORI:
				res = in.Op2 ^ in.Op3
			case in.Code == ins.LEAL:
				res = res | (in.Op2<<8|in.Op3)<<0
			case in.Code == ins.LEAH:
				res = res | (in.Op2<<8|in.Op3)<<16
			}
			au.Outputs["out"] <- res
		default:
		}
	}

}
