package exe

import (
  //"github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type ArithmeticUnit struct {
  com* comp.Communicator
}

func (au *ArithmeticUnit) Init() {
  au.com = &comp.Communicator{}
  au.com.Init()
}

func (au *ArithmeticUnit) Data() interface{} {
  return ""
}

func (au *ArithmeticUnit) State() string {
  return ""
}

func (au *ArithmeticUnit) Cycle() {

}

func (au *ArithmeticUnit) Communicator() * comp.Communicator {
  return au.com
}
