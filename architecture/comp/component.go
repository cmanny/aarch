package comp

import (
  "container/list"
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

func (c* Communicator) In(in string) chan interface {} {
  return c.Inputs[in]
}

func (c* Communicator) Out(out string) chan interface{} {
  return c.Outputs[out]
}
