package comp

import (
  "container/list"
  "fmt"
)

var comps *list.List

func Init() {
  comps = list.New()
}

func Add(c* Component) {
  comps.PushBack(c)
}

/* All components must satisfy the component functions */
type Component interface {
  Init()
  Data() interface{}
  State() string
  Cycle() /* advance one cycle */
}

/* All components need communicators */

func NewCommunicator() *Communicator {
  c := &Communicator{}
  c.InitComms()
  return c
}

type Communicator struct {
  Inputs map[string]chan interface{}
  Outputs map[string]chan interface{}
}

func (c* Communicator) InitComms() {
  c.Inputs = make(map[string]chan interface{})
  c.Outputs = make(map[string]chan interface{})
}

func (c* Communicator) In(p interface{}, in string) chan interface {} {
  fmt.Printf("%p \n", p)
  return c.Inputs[in]
}

func (c* Communicator) Out(p interface{}, out string) chan interface{} {
  fmt.Printf("%p \n", p)
  return c.Outputs[out]
}
