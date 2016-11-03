package comp

import (
  "fmt"
)

/* Component channel keys */
const (
  MEM_IN_1 = iota
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
  PIPE_LOGIC_IN
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

func Join(a Component, b Component, chanId int) {

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
  Inputs  map[string]chan interface{}
  Outputs map[string]chan interface{}
}

func (c *Communicator) InitComms() {
  c.Inputs = make(map[string]chan interface{})
  c.Outputs = make(map[string]chan interface{})
}

func (c *Communicator) In(ct Component, in string) chan interface{} {
  fmt.Printf("%p \n", &ct)

  return c.Inputs[in]
}

func (c *Communicator) Out(ct Component, out string) chan interface{} {
  fmt.Printf("%p \n", &ct)
  return c.Outputs[out]
}
