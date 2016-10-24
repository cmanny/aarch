package exe

import (
  "github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type MemoryUnit struct {
  com* Communicator
}

func (mu *Memory) Init() {
  mu.com = &Communicator{}
}

func (mu *Memory) Data() interface{} {
  return ""
}

func (mu *Memory) State() string {
  return ""
}

func (mu *Memory) Cycle() {

}

func (mu *Memory) Communicator() *Communicator {
  return mu.com
}
