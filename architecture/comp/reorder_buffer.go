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

type PhysicalRegister struct {
  Value int
  Valid bool
}

type ReorderBuffer struct {
  Communicator
  is *ins.InstructionSet

  buffer *list.List
  decoded *list.List
  freeNames *list.List

  rename [8][3]int
  physical [40]PhysicalRegister
}

func (rb* ReorderBuffer) Init(is *ins.InstructionSet) {
  rb.InitComms()
  rb.buffer = list.New()
  rb.decoded = list.New()
  rb.is = is

  rb.freeNames = list.New()
  //Fill up free renaming buffer
  rb.physical[0].Valid = true
  for i := 1; i < 40; i++ {
    rb.freeNames.PushBack(i)
    rb.physical[i].Valid = true
  }
  for i:= 0; i < 8; i ++{
    rb.rename[i] = [3]int{0, 0, 0}
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
    updateList := rb.Recv(CDB_RB_OUT).([]InsIn)

    for _, in := range updateList {
      if 0 < in.Tag && in.Tag < 40 {
        val, _ := rb.is.InsIdDecode(in.Code)
        fmt.Println(in)
        if val.Ins_type != ins.TYPE_CONTROL {
          rb.UpdatePhysicalRegister(in.Tag, in.Result, true)
          if rb.rename[in.RawOp1][0] != in.Tag {
            rb.freeNames.PushBack(in.Tag)
          }
        }
      }
    }

    //fmt.Println("re")

    entryList := list.New()
    rsList := list.New()

    for _, in := range decoded {
      //For each instruction arriving, give a buffer entry and tag

      in.Tag = rb.freeNames.Remove(rb.freeNames.Front()).(int)
      val, _ := rb.is.InsIdDecode(in.Code)
      if val.Ins_type != ins.TYPE_CONTROL {
        rb.UpdatePhysicalRegister(in.Tag, -1, false)
      }
      entry := BufferEntry{}
      entry.state = RB_ISSUED
      entry.in = in
      //fmt.Println(in)

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
      //fmt.Println(next.Value)
      rb.TagDeps(next.Value.(InsIn), next.Next())
      next = next.Next()
    }
    rb.decoded = rsList
  }
}

func (rb *ReorderBuffer) UpdateNamingTable(tagger InsIn, tag int) {
  rb.rename[tagger.RawOp1] = [3]int{tag, tag, tag}
}

// It is assumed that when a physical register is fulfilled, it is then valid
func (rb *ReorderBuffer) UpdatePhysicalRegister(tag int, value int, validity bool) {
  rb.physical[tag].Value = value
  rb.physical[tag].Valid = validity
}



func (rb *ReorderBuffer) UpdateByTables(in InsIn) InsIn {
  op3type, op2type := -1, -1
  if val, err := rb.is.InsIdDecode(in.Code); err == nil {
    op2type = val.Op2
    op3type = val.Op3
  }

  in.Op1 = int(int8(in.RawOp1))
  //Update by the latest known values
  if op2type == ins.OP_REGISTER {
    in.Op2Tag = rb.rename[in.RawOp2][0]
    in.Op2 = rb.physical[in.Op2Tag].Value
    in.Op2Valid = rb.physical[in.Op2Tag].Valid
  }
  if op2type == ins.OP_CONSTANT_8 {
    in.Op2 = int(in.RawOp2)
    in.Op2Valid = true
  }
  if op3type == ins.OP_REGISTER {
    in.Op3Tag = rb.rename[in.RawOp3][0]
    in.Op3 = rb.physical[in.Op3Tag].Value
    in.Op3Valid = rb.physical[in.Op3Tag].Valid
  }
  if op3type == ins.OP_CONSTANT_8 {
    in.Op3 = int(in.RawOp3)
    in.Op3Valid = true
  }
  if op2type == ins.OP_EMPTY {
    in.Op2Valid = true
  }
  if op3type == ins.OP_EMPTY {
    in.Op3Valid = true
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
    insInfo, _ := rb.is.InsIdDecode(updatedIns.Code)
    op2Type := insInfo.Op2
    op3Type := insInfo.Op3
    if updatedIns.Tag == tagger.Tag {
      break
    }
    if regId == updatedIns.RawOp2 && op2Type == ins.OP_REGISTER {
      updatedIns.Op2Tag = tag
      updatedIns.Op2Valid = false
    }
    if regId == updatedIns.RawOp3 && op3Type == ins.OP_REGISTER {
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
