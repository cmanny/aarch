package architecture

import (
  "fmt"
)

type Processor struct {
  clockSpeed int
  numExUnits int
  printDebug bool

  ip int

  rf RegisterFile
  is* InstructionSet
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


func (p* Processor) Init(is* InstructionSet) {
  p.is = is
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
}
