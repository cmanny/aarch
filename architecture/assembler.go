package architecture

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Assembler struct {
  inFile string
  is* InstructionSet
}

func (as* Assembler) Init(is* InstructionSet) {
  as.is = is
}

func (a* Assembler) LoadFile(fileName string) {
  f, err := os.Open(fileName)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    line := scanner.Text()
    words := strings.Fields(line)
    fmt.Println(words, len(words))
  }

}
