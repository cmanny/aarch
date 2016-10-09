package main

import (
  //"fmt"

  "github.com/cmanny/aarch/architecture"
)

/**
  In the main function we create the highest level parts of the system and
  connect them together. The assembler may be used to compile the written
  programs into bytecode which may then be placed into the memory component.

  Then, the processor can be run from an arbitrary chosen instruction pointer.

  The web frontend may also be enabled allowing the state of all components and
  channels to be automatically visualised by JS library.
**/

func main() {
  p := &architecture.Processor{}
  mem := &architecture.Memory{}
  as := &architecture.Assembler{}

  p.Debug(true)
  p.Run()
}
