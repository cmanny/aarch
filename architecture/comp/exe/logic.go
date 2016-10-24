package exe

import (
  //"github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type LogicUnit struct {
  com* comp.Communicator
}

func (lu *LogicUnit) Init() {
  lu.com = &comp.Communicator{}
  lu.com.Init()
}

func (lu *LogicUnit) Data() interface{} {
  return ""
}

func (lu *LogicUnit) State() string {
  return ""
}

func (lu *LogicUnit) Cycle() {

}

func (lu *LogicUnit) Communicator() *comp.Communicator {
  return lu.com
}
