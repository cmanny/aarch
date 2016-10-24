package comp

import (
    "bytes"
    "fmt"
)

type RegisterFile struct {
  regs [32] int
  Flags int

  com* Communicator
}

func (rf *RegisterFile) Init() {
  rf.com = &Communicator{}
}

func (rf *RegisterFile) Data() interface{} {
  return ""
}

func (rf *RegisterFile) State() string {
  return ""
}

func (rf *RegisterFile) Cycle() {

}

func (rf* RegisterFile) Contents() {
  var buffer bytes.Buffer
  for i := 0; i < 32; i++ {
    buffer.WriteString("r" + fmt.Sprintf("%d", i) + ": " + fmt.Sprintf("%d", rf.regs[i]) + "\n")
  }
  fmt.Println(buffer.String())
}
