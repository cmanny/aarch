package comp

/* All components must satisfy the component functions */
type Component interface {
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
