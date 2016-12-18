package comp

import (
  "fmt"
  "container/list"
  "github.com/cmanny/aarch/architecture/ins"
)

const (
  RB_ISSUED = iota
  RB_EXECUTING
  RB_FINISHED
)

type BufferEntry struct {
  in InsIn
  state int
}

type ReorderBuffer struct {
  Communicator
  is *ins.InstructionSet

  buffer *list.List
  decoded *list.List
  freeNames *list.List

  rename [8][3]int
  physical [40]int
}

func (rb* ReorderBuffer) Init(is *ins.InstructionSet) {
  rb.InitComms()
  rb.buffer = list.New()
  rb.decoded = list.New()
  rb.is = is

  rb.freeNames = list.New()
  //Fill up free renaming buffer
  for i := 0; i < 64; i++ {
    rb.freeNames.PushBack(i)
  }


}

func (rb* ReorderBuffer) Data() interface{} {
  return ""
}

func (rb* ReorderBuffer) State() string {
  return ""
}

func (rb* ReorderBuffer) Cycle() {
  for {
    rb.Recv(CYCLE)
    rb.Send(PIPE_RS_IN, rb.decoded)
    decoded := rb.Recv(PIPE_DECODE_OUT).([]InsIn)

    entryList := list.New()
    rsList := list.New()

    for _, in := range decoded {
      //For each instruction arriving, give a buffer entry and tag
      in.Tag = rb.freeNames.Remove(rb.freeNames.Front()).(int)
      entry := BufferEntry{}
      entry.state = RB_ISSUED
      entry.in = in
      fmt.Println(in)

      entryList.PushBack(entry)
      rsList.PushBack(in)
    }
    rb.buffer.PushBackList(entryList)
    //Now that the ROB has been updated, apply tags to dependent instructions
    //First revalidate with any new changes
    next := rsList.Front()
    for next != nil {
      next.Value = rb.UpdateByTables(next.Value.(InsIn))
      next = next.Next()
    }
    //Then tag any new dependencies
    next = rsList.Front()
    for next != nil {
      rb.TagDeps(next.Value.(InsIn), next.Next())
      next = next.Next()
    }
    rb.decoded = rsList
  }
}

func (rb *ReorderBuffer) UpdateNamingTable(tagger InsIn, tag int) {
  rb.rename[tagger.RawOp1] = [3]int{tag, tag, tag}
}

func (rb *ReorderBuffer) UpdatePhysicalRegister(tag int, value int) {
  rb.physical[tag] = value
}

func (rb *ReorderBuffer) UpdateByTables(in InsIn) InsIn {
  op3type, op2type := -1, -1
  if val, err := rb.is.InsIdDecode(in.Code); err == nil {
    op2type = val.Op2
    op3type = val.Op3
  }
  //Update by the latest known values
  if op2type == ins.OP_REGISTER {
    in.Op2Tag = rb.rename[in.RawOp2][0]
    in.Op2 = rb.physical[in.Op2Tag]
  }
  if op3type == ins.OP_REGISTER {
    in.Op3Tag = rb.rename[in.RawOp3][0]
    in.Op3 = rb.physical[in.Op3Tag]
  }
  return in
}


//Instructions that need to prevent hazards use this
func (rb *ReorderBuffer) TagDeps(tagger InsIn, next *list.Element) {
  regId := tagger.RawOp1
  tag := tagger.Tag
  if val, err := rb.is.InsIdDecode(tagger.Code);
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
    next = next.Next()
  }

  //Update the naming table, it is not the responsibility of this function to make this
  //A latest valid name
  rb.UpdateNamingTable(tagger, tag)
}
