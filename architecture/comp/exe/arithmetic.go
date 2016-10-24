package exe

import (
  "github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type ArithmeticUnit struct {

  Init()
  Data() interface{}
  State() string
  Cycle()
  Communicator() *Communicator
}

func (au *ArithmeticUnit) Init() {

}

func (au *ArithmeticUnit) Data() interface{} {

}

func (au *ArithmeticUnit) State() string {

}

func (au *ArithmeticUnit) Cycle() {

}

func (au *ArithmeticUnit) Communicator() *Communicator {

}
