package architecture

import (
  "fmt"
)

type Processor struct {
  clockSpeed int
  numExUnits int
  printDebug bool
  rf RegisterFile
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


func (p* Processor) Debug(toggle bool) {
  p.printDebug = toggle
}


func (p* Processor) Run() {
  p.preRun()
  fmt.Println("Processor beginning")
}
