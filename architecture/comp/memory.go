package comp

type Memory struct {
  bytes [4096] byte
  com *Communicator

}


func (mu *Memory) Init() {
  mu.com = &Communicator{}
  mu.com.Init()

  mu.com.Inputs["in1"] = make(chan interface{}, 1)
  mu.com.Outputs["out1"] = make(chan interface{}, 1)
}

func (mu *Memory) Data() interface{} {
  return ""
}

func (mu *Memory) State() string {
  return ""
}

func (mu *Memory) Cycle() {
  index := (<-mu.com.Inputs["in1"]).(int)
  mu.com.Outputs["out1"] <- mu.bytes[index: index + 4]
}

func (m *Memory) Communicator() *Communicator {
  return m.com
}

func (m* Memory) Fill(bytes []byte, index int) {
  for i := 0; i < len(bytes); i++ {
    m.bytes[i + index] = bytes[i]
  }
}
