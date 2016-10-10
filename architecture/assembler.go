package architecture

import (
  "bufio"
  "os"
  "strings"
  "fmt"
  "encoding/binary"
  "strconv"
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
      id, err := as.is.InsStrDecode(words[0])
      if err != nil {
        return nil, err
      }
      info, err := as.is.InsIdDecode(id)
      bytes = append(bytes, id)

      for i, op := range []int{info.Op1, info.Op2, info.Op3} {
        switch op {
          case  OP_REGISTER:
            regId, err := as.is.RegStrDecode(words[i + 1])
            if err != nil {
              return nil, err
              fmt.Println(regId)
            }
          case  OP_CONSTANT:
            bs := make([]byte, 4)
            val, err := strconv.Atoi(words[i + 1])
            if err != nil {
              return nil, err
            }
            binary.LittleEndian.PutUint32(bs, uint32(val))
            /* load constant into bytes */
          case  OP_ADDRESS:
            /* load address into bytes */
          case  OP_EMPTY:
            /* nothing */
          default:
        }
      }

    }
  }

  fmt.Println(as.labelIndex)
  fmt.Println(bytes)

  return bytes, err
}
