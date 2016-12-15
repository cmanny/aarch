package comp

//"github.com/cmanny/aarch/architecture/ins"

type MemoryUnit struct {
  Communicator
}

func (mu *MemoryUnit) Init() {
  mu.InitComms()
}

func (mu *MemoryUnit) Data() interface{} {
  return ""
}

func (mu *MemoryUnit) State() string {
  return ""
}

func (mu *MemoryUnit) Cycle() {

}
