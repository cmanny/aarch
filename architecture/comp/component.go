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
  Communicator() *Communicator /* retrieve the communicator */
}

/* All components need communicators */
type Communicator struct {
  Inputs map[string]chan interface{}
  Outputs map[string]chan interface{}
}
