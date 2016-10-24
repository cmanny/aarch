package exe

import (
  //"github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

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

}
