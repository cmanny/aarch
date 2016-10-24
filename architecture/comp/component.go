package comp

import (
	"fmt"
)

var Comps []Component

func Init() {
	Comps = make([]Component, 0)
}

func AddAll(cs ...Component) {
	for _, c := range cs {
		Comps = append(Comps, c)
	}
}

/* All components must satisfy the component functions */
type Component interface {
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
