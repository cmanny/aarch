package architecture

import (
  "fmt"

  "github.com/cmanny/aarch/architecture/comp"
  "github.com/cmanny/aarch/architecture/ins"
  //"bufio"
)

type Processor struct {
  comp.Communicator
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

  rs *comp.ReservationStation

  cu *comp.ControlUnit
  au1 *comp.ArithmeticUnit
  au2 *comp.ArithmeticUnit
  mu *comp.MemoryUnit

  bp *comp.BranchPredictor
  rb *comp.ReorderBuffer

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

func (p *Processor) execute(in *comp.InsIn) {
  if in == nil {
    return
  }

}

/**
  Public methods of Processor
**/

func (p *Processor) Init(is *ins.InstructionSet, mem *comp.Memory, cycle chan int) {
  comp.Init()
  p.InitComms()
  p.is = is
  p.cycle = cycle

  /* Init all sub components */
  p.mem = mem

  /* Fetch and decode init */
  p.fu = &comp.Fetch{}
  p.fu.Init()
  p.du = &comp.Decode{}
  p.du.Init()

  /* RS */

  p.rs = &comp.ReservationStation{}
  p.rs.Init()

  /* CU */

  p.cu = &comp.ControlUnit{}
  p.cu.Init()

  /* AU1 */

  p.au1 = &comp.ArithmeticUnit{}
  p.au1.Init()

  /* AU2 */

  p.au2 = &comp.ArithmeticUnit{}
  p.au2.Init()

  /* MU  */

  p.mu = &comp.MemoryUnit{}
  p.mu.Init()




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
      Name: "ArithmeticUnit1",
      Obj:  p.au1,
    },
    &comp.CompWrapper{
      Name: "ArithmeticUnit2",
      Obj:  p.au2,
    },
    &comp.CompWrapper{
      Name: "MemoryUnit",
      Obj:  p.mu,
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

  for _, c := range comp.Comps {
    comp.Join(p, c.Obj.(comp.Communicatizer), comp.CYCLE, 1)
    go c.Obj.(comp.Componentizer).Cycle()
  }

  for {
    for _, c := range comp.Comps {
      c.Obj.(comp.Communicatizer).Send(comp.CYCLE, 1)
    }

    if p.exit {
      return
    }
  }
}
