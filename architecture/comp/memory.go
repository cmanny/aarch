package comp

const (
  MEM_READ = iota
  MEM_WRITE
  MEM_RESET
)

type MemOp struct {
  Op int
  Addr int
  Len int
  Data []byte
}

type Memory struct {
  Communicator
  bytes [4096] byte
  opsPerCycle int
}

func (mu *Memory) Init() {
  mu.InitComms()
  mu.opsPerCycle = 2
}

func (mu *Memory) Data() interface{} {
  return ""
}

func (mu *Memory) State() string {
  return ""
}

func (mu *Memory) Cycle() {
  chans := [][]int{
    []int{MEM_IN_1, MEM_OUT_1},
    []int{MEM_IN_2, MEM_OUT_2},
    []int{MEM_IN_3, MEM_OUT_3},
  }

  for {
    mu.Recv(CYCLE)

    for i := 0; i < 5; i++ {
      for _, chanPair := range chans {
        ok, memOpIntf := mu.AsyncRecv(chanPair[0])
        if ok {
          memOp := memOpIntf.(MemOp)
          switch memOp.Op {
          case MEM_READ:
            mu.Send(chanPair[1], mu.bytes[memOp.Addr : memOp.Addr + memOp.Len])
          case MEM_WRITE:
            for j := 0; j < memOp.Len; j++ {
              mu.bytes[memOp.Addr + j] = memOp.Data[j]
            }
          }
        }
      }
    }
  }
}

func (m *Memory) Fill(bytes []byte, index int) {
  for i := 0; i < len(bytes); i++ {
    m.bytes[i+index] = bytes[i]
  }
}
