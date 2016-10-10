package main

import (
  "fmt"
  "log"
  "flag"

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

  fileNamePtr := flag.String("prog", "fib.gp", "file to assemble and run")
  runWebPtr := flag.Bool("web", true, "bool to run web front end")

  flag.Parse()

  /* Init top level components */

  is  := &architecture.InstructionSet{}
  is.Init()
  p   := &architecture.Processor{}
  p.Init(is)
  mem := &architecture.Memory{}
  mem.Init()
  as  := &architecture.Assembler{}
  as.Init(is)

  /* Begin */

  bytes, err := as.AssembleFile(*fileNamePtr)

  if err != nil {
    log.Fatal(err)
  }

  /* Fill memory with compilde bytes */

  mem.Fill(bytes, 0)




  fmt.Errorf(*fileNamePtr, *runWebPtr, *mem, *as)




  p.Debug(true)
  p.Run()
}
