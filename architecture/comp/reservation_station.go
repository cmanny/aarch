package comp

type ReservationStation struct {
  Communicator


}

func (rs* ReservationStation) Init() {
  rs.InitComms()
}

func (rs* ReservationStation) Data() interface{} {
  return ""
}

func (rs* ReservationStation) State() string {
  return ""
}

func (rs* ReservationStation) Cycle() {
  for {
    rs.Recv(CYCLE)
    /* Send out all shelving buffers */

    /* Receive bypass and rob pass through */


  }
}
