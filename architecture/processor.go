package architecture

import (
  "fmt"

  "github.com/cmanny/aarch/architecture/comp"
  "github.com/cmanny/aarch/architecture/comp/exe"
  "github.com/cmanny/aarch/architecture/ins"
  //"bufio"
)

type Processor struct {
  clockSpeed int
  numExUnits int
  printDebug bool
  exit       bool

  cycle chan int

  ip int

  is *ins.InstructionSet

  /* Components */

  mem *comp.Memory

  fu *comp.Fetch
  du *comp.Decode

  cu *exe.ControlUnit
  au *exe.ArithmeticUnit
  lu *exe.LogicUnit

  rf *comp.RegisterFile
}

/**
  Private methods of Processor
**/

func (p *Processor) preRun() {
  if p.printDebug {
    fmt.Println("Debug ON")
  }
}

func (p *Processor) execute(in *exe.InsIn) {
  if in == nil {
    return
  }

}

/**
  Public methods of Processor
**/

func (p *Processor) Init(is *ins.InstructionSet, mem *comp.Memory, cycle chan int) {
  comp.Init()
  p.is = is
  p.cycle = cycle

  /* Init all sub components */
  p.mem = mem

  p.fu = &comp.Fetch{}
  p.du = &comp.Decode{}

  p.cu = &exe.ControlUnit{}
  p.cu.Init()

  p.au = &exe.ArithmeticUnit{}
  p.au.Init()

  p.lu = &exe.LogicUnit{}
  p.lu.Init()

  comp.AddAll(
    &comp.CompWrapper{
      Name: "RAM",
      Obj:  p.mem,
    },
    &comp.CompWrapper{
      Name: "Fetch",
      Obj:  p.fu,
    },
    &comp.CompWrapper{
      Name: "Decode",
      Obj:  p.du,
    },
    &comp.CompWrapper{
      Name: "ControlUnit",
      Obj:  p.cu,
    },
    &comp.CompWrapper{
      Name: "ArithmeticUnit",
      Obj:  p.au,
    },
    &comp.CompWrapper{
      Name: "LogicUnit",
      Obj:  p.lu,
    },
  )
}

func (p *Processor) Data() interface{} {
  return 1
}

func (p *Processor) State() string {
  return ""
}

func (p *Processor) Cycle() {

}

func (p *Processor) Debug(toggle bool) {
  p.printDebug = toggle
}

func (p *Processor) SetIP(ip int) {
  p.ip = ip
}

func (p *Processor) Run() {
  p.preRun()
  fmt.Println("Processor beginning")

  for {
    for _, c := range comp.Comps {
      go c.Obj.Cycle()
    }

    if p.exit {
      return
    }
  }
}
