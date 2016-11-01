package comp

type Memory struct {
  Communicator
  bytes [4096]byte
}

func (mu *Memory) Init() {
  mu.InitComms()
  mu.Inputs["in1"] = make(chan interface{}, 1)
  mu.Outputs["out1"] = make(chan interface{}, 1)
}

func (mu *Memory) Data() interface{} {
  return ""
}

func (mu *Memory) State() string {
  return ""
}

func (mu *Memory) Cycle() {
  index := (<-mu.Inputs["in1"]).(int)
  mu.Outputs["out1"] <- mu.bytes[index : index+4]
}

func (m *Memory) Fill(bytes []byte, index int) {
  for i := 0; i < len(bytes); i++ {
    m.bytes[i+index] = bytes[i]
  }
}
