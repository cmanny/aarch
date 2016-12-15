package comp

type BranchPredictor struct {
  Communicator
}

func (bp* BranchPredictor) Init() {
  bp.InitComms()
}

func (bp* BranchPredictor) Data() interface{} {
  return ""
}

func (bp* BranchPredictor) State() string {
  return ""
}

func (bp* BranchPredictor) Cycle() {
}
