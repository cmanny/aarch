package comp

import (
	"fmt"
)

var Comps []*CompWrapper

func Init() {
	Comps = make([]*CompWrapper, 0)
}

func AddAll(cs ...*CompWrapper) {
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

type CompWrapper struct {
	Name string
	Obj  Component
}

type NullComponent struct{}

func (n *NullComponent) Data() interface{} { return 0 }
func (n *NullComponent) State() string     { return "" }
func (n *NullComponent) Cycle()            {}

/* All components need communicators */

type Communicator struct {
	Inputs    map[string]chan interface{}
	Outputs   map[string]chan interface{}
	MustCycle bool
}

func (c *Communicator) InitComms() {
	c.Inputs = make(map[string]chan interface{})
	c.Outputs = make(map[string]chan interface{})
}

func (c *Communicator) In(ct Component, in string) chan interface{} {
	c.MustCycle = true
	fmt.Printf("%p \n", &ct)

	return c.Inputs[in]
}

func (c *Communicator) Out(ct Component, out string) chan interface{} {
	fmt.Printf("%p \n", &ct)
	return c.Outputs[out]
}
