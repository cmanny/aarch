package comp

import (
  "bytes"
  "fmt"
)

const (
  REG_READ = iota
  REG_WRITE
  REG_INVALIDATE
  REG_CLEAR
)

type RegData struct {
  Id int
  Value int
  Tag int
  Valid bool
}

type RegOp struct {
  Op int
  Data RegData
}

type RegisterFile struct {
  Communicator
  regs  [32] RegData
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
    for _, chanPair := range chans {
      ok, regOpsIntf := rf.AsyncRecv(chanPair[0])
      if ok {
        regOps := regOpsIntf.([]RegOp)
        returnBuf := make([]RegData, 0)
        for _, regOp := range regOps {
          switch regOp.Op {
          case REG_READ:
            returnBuf = append(returnBuf, rf.regs[regOp.Data.Id])
          case REG_WRITE:
            rf.regs[regOp.Data.Id].Value = regOp.Data.Value
            if rf.regs[regOp.Data.Id].Tag == regOp.Data.Tag {
              rf.regs[regOp.Data.Id].Valid = true
            }

          case REG_INVALIDATE:
            rf.regs[regOp.Data.Id].Valid = false
            rf.regs[regOp.Data.Id].Tag = regOp.Data.Tag
          }
        }
        if len(returnBuf) > 0 {
          rf.Send(chanPair[1], returnBuf)
        }
      }
    }
  }
}

func (rf *RegisterFile) Contents() {
  var buffer bytes.Buffer
  for i := 0; i < 32; i++ {
    buffer.WriteString("r" + fmt.Sprintf("%d", i) + ": " + fmt.Sprintf("%d", rf.regs[i].Value) + "\n")
  }
  fmt.Println(buffer.String())
}
