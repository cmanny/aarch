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

func (au *Memory) Init() {
  au.com = &Communicator{}
}

func (au *Memory) Data() interface{} {
  return ""
}

func (au *Memory) State() string {
  return ""
}

func (au *Memory) Cycle() {

}

func (au *Memory) Communicator() *Communicator {
  return au.com
}
