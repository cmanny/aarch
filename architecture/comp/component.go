package comp


/* Component channel keys */
const (
  CYCLE = iota

  MEM_IN_1
  MEM_IN_2
  MEM_IN_3
  MEM_OUT_1
  MEM_OUT_2
  MEM_OUT_3

  REG_IN_1
  REG_IN_2
  REG_IN_3
  REG_OUT_1
  REG_OUT_2
  REG_OUT_3

  PIPE_DECODE_IN
  PIPE_ARITH_IN
  PIPE_CONTROL_IN
  PIPE_MEMORY_IN

  PIPE_DECODE_OUT
  PIPE_ARITH_OUT
  PIPE_CONTROL_OUT
  PIPE_MEMORY_OUT
)

var Comps []*CompWrapper
var Joins []Edge

func Init() {
  Comps = make([]*CompWrapper, 0)
}

func AddAll(cs ...*CompWrapper) {
  for _, c := range cs {
    Comps = append(Comps, c)
  }
}

func Join(a *Communicator, b *Communicator, chanId int, bufSize int) {
  chanRef := make(chan interface{}, bufSize)
  a.SetChan(chanId, chanRef)
  b.SetChan(chanId, chanRef)

}

/* All components must satisfy the component functions */
type Component interface {
  Data() interface{}
  State() string
  Cycle() /* advance one cycle */
}


type Edge struct{
  A Component
  B Component
  data interface{}
}

type CompWrapper struct {
  Name  string
  Obj   Component
  Shape string
  Size  int
  Color string
}

type NullComponent struct{}

func (n *NullComponent) Data() interface{} { return 0 }
func (n *NullComponent) State() string     { return "" }
func (n *NullComponent) Cycle()            {}

/* All components need communicators */

type Communicator struct {
  chans  map[int]chan interface{}
  recvd  map[int]interface{}
}

func (c *Communicator) InitComms() {
  c.chans = make(map[int]chan interface{})
  c.recvd = make(map[int]interface{})
}

func (c *Communicator) Send(chanId int, data interface{}) {
  c.chans[chanId] <- data
}

func (c *Communicator) Recv(chanId int) interface{} {
  c.recvd[chanId] = <- c.chans[chanId]
  return c.recvd[chanId]
}

func (c *Communicator) SetChan(chanId int, chanRef chan interface{}) {
  c.chans[chanId] = chanRef
}
