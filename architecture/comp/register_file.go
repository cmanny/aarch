package comp

import (
  "bytes"
  "fmt"
)

const (
  REG_READ = iota
  REG_WRITE
  REG_CLEAR
)

type RegOp struct {
  Op int
  Id int
  Len int
  Data [4]byte
}

type RegisterFile struct {
  Communicator
  regs  [32][4]byte
  Flags int
}

func (rf *RegisterFile) Init() {
  rf.InitComms()
}

func (rf *RegisterFile) Data() interface{} {
  return ""
}

func (rf *RegisterFile) State() string {
  return ""
}

func (rf *RegisterFile) Cycle() {
  chans := [][]int{
    []int{REG_IN_1, REG_OUT_1},
    []int{REG_IN_2, REG_OUT_2},
    []int{REG_IN_3, REG_OUT_3},
  }

  for {
    rf.Recv(CYCLE)

    for i := 0; i < 5; i++ {
      for _, chanPair := range chans {
        ok, regOpIntf := rf.AsyncRecv(chanPair[0])
        if ok {
          regOp := regOpIntf.(RegOp)
          switch regOp.Op {
          case REG_READ:
            rf.Send(chanPair[1], rf.regs[regOp.Id])
          case REG_WRITE:
            rf.regs[regOp.Id] = regOp.Data
          }
        }
      }
    }
  }
}

func (rf *RegisterFile) Contents() {
  var buffer bytes.Buffer
  for i := 0; i < 32; i++ {
    buffer.WriteString("r" + fmt.Sprintf("%d", i) + ": " + fmt.Sprintf("%d", rf.regs[i]) + "\n")
  }
  fmt.Println(buffer.String())
}
