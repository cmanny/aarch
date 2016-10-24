package exe

import (
  "github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type LogicUnit struct {
  com* Communicator
}

func (lu *LogicUnit) Init() {
  lu.com = &Communicator{}
}

func (lu *LogicUnit) Data() interface{} {
  return ""
}

func (lu *LogicUnit) State() string {
  return ""
}

func (lu *LogicUnit) Cycle() {

}

func (lu *LogicUnit) Communicator() *Communicator {
  return lu.com
}
