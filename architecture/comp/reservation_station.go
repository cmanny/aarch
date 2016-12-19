package comp

import (
  "container/list"

  "github.com/cmanny/aarch/architecture/ins"
)

type Shelf struct {
  In InsIn
  Type int
  Filled bool
  ChanId int
}

type ReservationStation struct {
  Communicator
  is* ins.InstructionSet

  queue *list.List


  au1Shelf Shelf
  au2Shelf Shelf
  cuShelf Shelf
  muShelf Shelf

  shelves [4]*Shelf

}

func (rs* ReservationStation) Init(is *ins.InstructionSet) {
  rs.InitComms()
  rs.queue = list.New()
  rs.is = is

  rs.au1Shelf = Shelf{InsIn{}, ins.TYPE_ARITH, false, PIPE_ARITH_IN_1}
  rs.au2Shelf = Shelf{InsIn{}, ins.TYPE_ARITH, false, PIPE_ARITH_IN_2}
  rs.cuShelf = Shelf{InsIn{}, ins.TYPE_CONTROL, false, PIPE_CONTROL_IN}
  rs.muShelf = Shelf{InsIn{}, ins.TYPE_MOVE, false, PIPE_MEMORY_IN}
  rs.shelves = [4]*Shelf{&rs.au1Shelf, &rs.au2Shelf, &rs.cuShelf, &rs.muShelf}

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
    for _, s := range rs.shelves {
      in := InsIn{}
      in.Tag = -1
      if s.Filled {
        in = s.In
        s.Filled = false
      }
      rs.Send(s.ChanId, in)
    }

    // Get new instructions from ROB
    entryList := rs.Recv(PIPE_RS_IN).(*list.List)
    //fmt.Println(entryList)
    updateList := rs.Recv(CDB_RS_OUT).([]InsIn)

    //Combine new entries into RS list
    rs.queue.PushBackList(entryList)
    //Resolve tags from CDB
    for _, in := range updateList {
      rs.ResolveTags(in, rs.queue.Front())
    }

    //Look up inthe list and fill up shelves
    next := rs.queue.Front()
    toRemove := make([]*list.Element, 0)
    for next != nil {
      in := next.Value.(InsIn)
      //fmt.Println(in)
      if in.Op2Valid && in.Op3Valid {
        //fmt.Println("Found valid")
        //We found a ready instruction
        val, err := rs.is.InsIdDecode(in.Code);
        if err != nil {
          continue
        }
        for _, s := range rs.shelves {
          if !s.Filled && s.Type == val.Ins_type {
            toRemove = append(toRemove, next)
            s.In = in
            s.Filled = true
            break
          }
        }
      }
      next = next.Next()
    }
    for _, r := range toRemove {
      rs.queue.Remove(r)
    }

  }
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
    //fmt.Println("Resolving tags")
    updatedIns := next.Value.(InsIn)
    if updatedIns.Op2Tag == tag {
      updatedIns.Op2 = result
      updatedIns.Op2Valid = true
    }
    if updatedIns.Op3Tag == tag {
      updatedIns.Op3 = result
      updatedIns.Op3Valid = true
    }
    next.Value = updatedIns
    next = next.Next()
  }
}
