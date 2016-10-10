package architecture

type Memory struct {
  Data [4096] byte
}

func (m* Memory) Init() {
  
}

func (m* Memory) Fill(bytes []byte, index int) {
  for i := 0; i < len(bytes); i++ {
    m.Data[i + index] = bytes[i]
  }
}
