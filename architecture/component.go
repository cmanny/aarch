package architecture

/* All components must satisfy the component functions */
type Component interface {
  Data() interface{}
  State() string
}
