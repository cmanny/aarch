package architecture

import (
  "bufio"
  "os"
  "strings"
  "fmt"
)

type Assembler struct {
  inFile string
  is* InstructionSet

  labelIndex map[string]int
}

func (as* Assembler) Init(is* InstructionSet) {
  as.is = is

  as.labelIndex = make(map[string]int)
}

func (as* Assembler) AssembleFile(fileName string) ([]byte, error) {
  f, err := os.Open(fileName)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  bytes := make([]byte, 0)

  scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

  bytes = append(bytes, 0)

  byteIndex := 0
  for scanner.Scan() {
    line := scanner.Text()
    words := strings.Fields(line)
    if len(words) == 0 {
      continue
    }
    fmt.Println(words)
    if words[0][0] == '.' {
      /* label */
      as.labelIndex[words[0][1:]] = byteIndex
    } else {
      /* normal instruction OR data*/

    }
  }

  fmt.Println(as.labelIndex)

  return bytes, err
}
