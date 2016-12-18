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

  rename map[byte][3]int //maps to one of 3 flow states
  physical map[int]int


  au1Shelf InsIn
  au2Shelf InsIn
  cuShelf InsIn
  muShelf InsIn

}

func (rs* ReservationStation) Init(is *ins.InstructionSet) {
  rs.InitComms()
  rs.queue = list.New()
  rs.is = is

  rs.rename = make(map[byte][3]int)
  rs.physical = make(map[int]int)
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

func (rs *ReservationStation) UpdateNamingTable(tagger InsIn, tag int) {
  rs.rename[tagger.RawOp1] = [3]int{tag, tag, tag}
}

func (rs *ReservationStation) UpdatePhysicalRegister(tag int, value int) {
  rs.physical[tag] = value
}


//Instructions that need to prevent hazards use this
func (rs *ReservationStation) TagDeps(tagger InsIn, next *list.Element) {
  regId := tagger.RawOp1
  tag := tagger.Tag
  if val, err := rs.is.InsIdDecode(tagger.Code);
    err != nil || val.Ins_type == ins.TYPE_CONTROL || val.Op1 == ins.OP_IND_ADDR {
    return
  }

  //Search up the chain for any tags to apply
  for next != nil {
    updatedIns := next.Value.(InsIn)
    if regId == updatedIns.RawOp2 {
      updatedIns.Op2Tag = tag
      updatedIns.Op2Valid = false
    }
    if regId == updatedIns.RawOp3 {
      updatedIns.Op3Tag = tag
      updatedIns.Op3Valid = false
    }
    next.Value = updatedIns
  }

  //Update the naming table, it is not the responsibility of this function to make this
  //A latest valid name
  rs.UpdateNamingTable(tagger, tag)
}

// Instructions that have just returned from the bypass will use this
func (rs *ReservationStation) ResolveTags(tagger InsIn, next *list.Element) {
  tag := tagger.Tag
  result := tagger.Result

  if val, err := rs.is.InsIdDecode(tagger.Code);
    err != nil || val.Ins_type == ins.TYPE_CONTROL || val.Op1 == ins.OP_IND_ADDR {
    return
  }

  //Search up the chain for any tags to resolve
  for next != nil {
    updatedIns := next.Value.(InsIn)
    if updatedIns.Op2Tag == tag {
      updatedIns.Op2 = result
      updatedIns.Op2Valid = true
    }
    if updatedIns.Op3Tag == tag {
      updatedIns.Op3 = tag
      updatedIns.Op3Valid = true
    }
    next.Value = updatedIns
  }
  // After resolving any tags, update the physical table
  rs.UpdatePhysicalRegister(tag, result)
}
