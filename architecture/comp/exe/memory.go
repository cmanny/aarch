package exe

//"github.com/cmanny/aarch/architecture/ins"
import "github.com/cmanny/aarch/architecture/comp"

type MemoryUnit struct {
	comp.Communicator
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
