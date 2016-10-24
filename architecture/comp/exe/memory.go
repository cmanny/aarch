package exe

import (
  //"github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type MemoryUnit struct {
  com* comp.Communicator
}

func (mu *MemoryUnit) Init() {
  mu.com = &comp.Communicator{}
  mu.com.Init()
}

func (mu *MemoryUnit) Data() interface{} {
  return ""
}

func (mu *MemoryUnit) State() string {
  return ""
}

func (mu *MemoryUnit) Cycle() {

}

func (mu *MemoryUnit) Communicator() *comp.Communicator {
  return mu.com
}
