package comp

type Memory struct {
  bytes [4096] byte
  com *Communicator

}


func (m *Memory) Init() {
  m.com = &Communicator{}
}

func (m *Memory) Data() interface{} {
  return ""
}

func (m *Memory) State() string {
  return ""
}

func (m *Memory) Cycle() {

}

func (m *Memory) Communicator() *Communicator {
  return m.com
}

func (m* Memory) Fill(bytes []byte, index int) {
  for i := 0; i < len(bytes); i++ {
    m.bytes[i + index] = bytes[i]
  }
}
