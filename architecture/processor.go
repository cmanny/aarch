package architecture

import (
  "fmt"
  "github.com/cmanny/aarch/architecture/comp"
  "github.com/cmanny/aarch/architecture/comp/exe"
  "github.com/cmanny/aarch/architecture/ins"
  "bufio"
  "os"
)

type Processor struct {
  clockSpeed int
  numExUnits int
  printDebug bool

  ip int

  is* ins.InstructionSet
  mem* comp.Memory
  cu* exe.ControlUnit
}


/**
  Private methods of Processor
**/


func (p* Processor) preRun() {
  if p.printDebug {
    fmt.Println("Debug ON")
  }
}

func (p* Processor) fetch() {
  fmt.Println("trying to fetch")
  bufio.NewReader(os.Stdin).ReadString('\n')
  p.mem.Communicator().Inputs["in1"] <- p.ip
  bytes := <- p.mem.Communicator().Outputs["out1"]
  fmt.Println(bytes)

  p.cu.Communicator().Inputs["ins"] <- bytes
  p.ip = (<- p.cu.Communicator().Outputs["ip"]).(int)


}

func (p* Processor) decode() {

}

func (p* Processor) execute() {
}

func (p* Processor) writeback() {

}
/**
  Public methods of Processor
**/


func (p* Processor) Init(is* ins.InstructionSet, mem* comp.Memory) {
  p.is = is
  p.mem = mem

  p.cu = &exe.ControlUnit{}
  p.cu.Init()
  p.ip = (<-p.cu.Communicator().Outputs["ip"]).(int)
}

func (p* Processor) Debug(toggle bool) {
  p.printDebug = toggle
}

func (p* Processor) SetIP(ip int) {
  p.ip = ip
}


func (p* Processor) Run() {
  p.preRun()
  fmt.Println("Processor beginning")



  for {
    go p.mem.Cycle()
    go p.cu.Cycle()
    p.fetch()
    p.decode()
    p.execute()
  }
}
