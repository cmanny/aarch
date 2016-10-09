package architecture

import (
    "bytes"
    "fmt"
)

type RegisterFile struct {
  Data [32] int
  Flags int
}

func (rf* RegisterFile) Contents() {
  var buffer bytes.Buffer
  for i := 0; i < 32; i++ {
    buffer.WriteString("r" + fmt.Sprintf("%d", i) + ": " + fmt.Sprintf("%d", rf.Data[i]) + "\n")
  }
  fmt.Println(buffer.String())
}
