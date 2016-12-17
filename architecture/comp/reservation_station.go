package comp

import (
  "fmt"
  "container/list"

  "github.com/cmanny/aarch/architecture/ins"
)

type ReservationStation struct {
  Communicator
  is* ins.InstructionSet

  queue *list.List


  au1Shelf InsIn
  au2Shelf InsIn
  cuShelf InsIn
  muShelf InsIn

}

func (rs* ReservationStation) Init(is *ins.InstructionSet) {
  rs.InitComms()
  rs.queue = list.New()
  rs.is = is
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

    // Get new instructions from ROB
    entryList := rs.Recv(PIPE_RS_IN).(*list.List)
    rs.queue.PushBackList(entryList)
    fmt.Println(rs.queue.Front())

    /* Send out all shelving buffers */

    /* Receive bypass and rob pass through */


  }
}
