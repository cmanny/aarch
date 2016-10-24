package exe

import (
  //"github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type LogicUnit struct {
  comp.Communicator
}

func (lu *LogicUnit) Init() {
  lu.InitComms()
}

func (lu *LogicUnit) Data() interface{} {
  return ""
}

func (lu *LogicUnit) State() string {
  return ""
}

func (lu *LogicUnit) Cycle() {

}
