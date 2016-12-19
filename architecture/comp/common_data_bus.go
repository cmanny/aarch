package comp

//import "fmt"

type CommonDataBus struct {
  Communicator
}

func (cdb* CommonDataBus) Init() {
  cdb.InitComms()
}

func (cdb* CommonDataBus) Data() interface{} {
  return ""
}

func (cdb* CommonDataBus) State() string {
  return ""
}

func (cdb* CommonDataBus) Cycle() {
  for {
    cdb.Recv(CYCLE)

    au1Out := cdb.Recv(PIPE_ARITH_OUT_1).(InsIn)
    au2Out := cdb.Recv(PIPE_ARITH_OUT_2).(InsIn)
    cuOut := cdb.Recv(PIPE_CONTROL_OUT).(InsIn)
    muOut := cdb.Recv(PIPE_MEMORY_OUT).(InsIn)


    array := []InsIn{au1Out, au2Out, cuOut, muOut}

    cdb.Send(CDB_RB_OUT, array)
    cdb.Send(CDB_RS_OUT, array)

  }
}
