package comp

import (
   //"fmt"
  // "reflect"
)

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

  PIPE_FETCH_IN
  PIPE_DECODE_IN
  PIPE_ARITH_IN
  PIPE_ARITH_IN_1
  PIPE_ARITH_IN_2
  PIPE_CONTROL_IN
  PIPE_MEMORY_IN
  PIPE_RS_IN

  PIPE_DECODE_OUT
  PIPE_ARITH_OUT
  PIPE_ARITH_OUT_1
  PIPE_ARITH_OUT_2
  PIPE_CONTROL_OUT
  PIPE_MEMORY_OUT

  CDB_RS_OUT
  CDB_RB_OUT
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

func Join(a Communicatizer, b Communicatizer, chanId int, bufSize int) {
  chanRef := make(chan interface{}, bufSize)
  a.SetChan(chanId, chanRef)
  b.SetChan(chanId, chanRef)

}

func JoinDifferent(a Communicatizer, aChanId int, b Communicatizer, bChanId int, bufSize int) {
  chanRef := make(chan interface{}, bufSize)
  a.SetChan(aChanId, chanRef)
  b.SetChan(bChanId, chanRef)

}

/* All components must satisfy the component functions */
type Componentizer interface {
  Data() interface{}
  State() string
  Cycle() /* advance one cycle */
}

type PipelineData struct {
  current interface{}
  next interface{}
}


type Edge struct{
  A Componentizer
  B Componentizer
  data interface{}
}

type CompWrapper struct {
  Name  string
  Obj   interface{}
  Shape string
  Size  int
  Color string
}

type NullComponent struct{}

func (n *NullComponent) Data() interface{} { return 0 }
func (n *NullComponent) State() string     { return "" }
func (n *NullComponent) Cycle()            {}

/* All components need communicators */

type Communicatizer interface {
  Send(chanId int, data interface{})
  AsyncSend(chanId int, data interface{}) bool
  Recv(chanId int) interface{}
  AsyncRecv(chanId int) (bool, interface{})

  SetChan(int, chan interface{})
}

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

func (c *Communicator) AsyncSend(chanId int, data interface{}) bool {
  select {
  case c.chans[chanId] <- data:
    return true
    default:
  }
  return false
}

func (c *Communicator) Recv(chanId int) interface{} {
  c.recvd[chanId] = <- c.chans[chanId]
  return c.recvd[chanId]
}


func (c *Communicator) AsyncRecv(chanId int) (bool, interface{}) {
  //fmt.Println(c.chans)
  select {
    case c.recvd[chanId] = <- c.chans[chanId]:

      return true, c.recvd[chanId]
    default:
  }
  return false, c.recvd[chanId]
}

func (c *Communicator) SetChan(chanId int, chanRef chan interface{}) {
  c.chans[chanId] = chanRef
}
