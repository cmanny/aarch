package comp

import (
  "fmt"
  "container/list"
)

type ReservationStation struct {
  Communicator

  queue *list.List


  au1Shelf InsIn
  au2Shelf InsIn
  cuShelf InsIn
  muShelf InsIn

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
    /* Send all shelved values */
    // rs.Send(PIPE_ARITH_IN_1, rs.au1Shelf)
    // rs.Send(PIPE_ARITH_IN_2, rs.au1Shelf)
    // rs.Send(PIPE_MEMORY_IN, rs.muShelf)
    // rs.Send(PIPE_CONTROL_IN, rs.cuShelf)

    in := rs.Recv(PIPE_RS_IN).(*list.List)
    if in.Len() > 0 {}
    get := in.Back()
    if get != nil {
      fmt.Println(get.Value)
    }
    //fmt.Println(in)

    /* Send out all shelving buffers */

    /* Receive bypass and rob pass through */


  }
}
