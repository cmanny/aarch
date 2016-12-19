package architecture

import (
  "fmt"
  "time"

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
  numCycles int

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

  cdb *comp.CommonDataBus

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
  p.numCycles = 0

  /* Init all sub components */
  p.mem = mem

  /* Fetch and decode init */
  p.fu = &comp.Fetch{}
  p.fu.Init(is)
  p.du = &comp.Decode{}
  p.du.Init()

  /* RS */

  p.rs = &comp.ReservationStation{}
  p.rs.Init(is)

  /* ROB */

  p.rb = &comp.ReorderBuffer{}
  p.rb.Init(is)

  /* BP */

  p.bp = &comp.BranchPredictor{}
  p.bp.Init()

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

  /* RF */

  p.rf = &comp.RegisterFile{}
  p.rf.Init()

  p.cdb = &comp.CommonDataBus{}
  p.cdb.Init()





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
      Name: "ReservationStation",
      Obj:  p.rs,
    },
    &comp.CompWrapper{
      Name: "ReorderBuffer",
      Obj:  p.rb,
    },
    &comp.CompWrapper{
      Name: "BranchPredictor",
      Obj:  p.bp,
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
    &comp.CompWrapper{
      Name: "CommonDataBus",
      Obj:  p.cdb,
    },
  )

  /* Joining components (important (do not forget))
    The joining design isn't the best but it suffices for now
    A better method would be much more explicit about channel identifiers
    and would guarantee some level of correctness
  */


  /* Mem to fetch */
  comp.Join(p.mem, p.fu, comp.MEM_IN_1, 1)
  comp.Join(p.mem, p.fu, comp.MEM_OUT_1, 1)

  /* Fetch to Decode */
  comp.Join(p.fu, p.du, comp.PIPE_DECODE_IN, 1)

  /* Decode to ReorderBuffer */
  comp.Join(p.du, p.rb, comp.PIPE_DECODE_OUT, 1)

  comp.Join(p.rb, p.rs, comp.PIPE_RS_IN, 1)

  comp.JoinDifferent(p.rs, comp.PIPE_ARITH_IN_1, p.au1, comp.PIPE_ARITH_IN, 1)
  comp.JoinDifferent(p.rs, comp.PIPE_ARITH_IN_2, p.au2, comp.PIPE_ARITH_IN, 1)
  comp.Join(p.rs, p.cu, comp.PIPE_CONTROL_IN, 1)
  comp.Join(p.rs, p.mu, comp.PIPE_MEMORY_IN, 1)

  comp.JoinDifferent(p.cdb, comp.PIPE_ARITH_OUT_1, p.au1, comp.PIPE_ARITH_OUT, 1)
  comp.JoinDifferent(p.cdb, comp.PIPE_ARITH_OUT_2, p.au2, comp.PIPE_ARITH_OUT, 1)
  comp.Join(p.cdb, p.cu, comp.PIPE_CONTROL_OUT, 1)
  comp.Join(p.cdb, p.mu, comp.PIPE_MEMORY_OUT, 1)

  comp.Join(p.cdb, p.rs, comp.CDB_RS_OUT, 1)
  comp.Join(p.cdb, p.rb, comp.CDB_RB_OUT, 1)

  comp.Join(p.fu, p.rb, comp.PIPE_FETCH_IN, 1)


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
    p.numCycles++
    for _, c := range comp.Comps {
      if !c.Obj.(comp.Communicatizer).AsyncSend(comp.CYCLE, 1) {
        //fmt.Println(fmt.Sprintf("%p did not receive\n", c.Obj))
      } else {
        //fmt.Println(fmt.Sprintf("%p did receive\n", c.Obj))
      }

    }
    time.Sleep(time.Millisecond * 100)

    if p.exit {
      return
    }
  }
}
