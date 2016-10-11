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

func (as* Assembler) parse(scanner* bufio.Scanner, labelOnly bool) ([]byte, error) {

  byteIndex := 0
  bytes := make([]byte, 0)
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
      byteIndex += 1

      for i, op := range []int{info.Op1, info.Op2, info.Op3} {
        /* First pass, only labels are resolved, so we only increment byte index */
        if labelOnly {
          if op == OP_ADDRESS_32 || op == OP_CONSTANT_32 {
            byteIndex += 4
          }else{
            byteIndex += 1
          }
          continue
        }
        /* Second pass we generate the bytes and add them to the return buffer */
        switch op {
          case  OP_REGISTER:
            regId, err := as.is.RegStrDecode(words[i + 1])
            if err != nil {
              return nil, err
            }
            bytes = append(bytes, regId)
          case OP_CONSTANT_32:
            bs := make([]byte, 4)
            val, err := strconv.Atoi(words[i + 1])
            if err != nil {
              return nil, err
            }
            binary.LittleEndian.PutUint32(bs, uint32(val))
            bytes = append(bytes, bs...)

          case OP_CONSTANT_8:
            /* load constant into bytes */
          case  OP_ADDRESS_8:
            bs := make([]byte, 4)
            val, err := strconv.Atoi(words[i + 1])
            if err != nil {
              return nil, err
            }
            binary.LittleEndian.PutUint32(bs, uint32(val))
            bytes = append(bytes, bs...)
            /* load address into bytes */
          case  OP_EMPTY:
            /* nothing */
          default:
        }
      }

    }
  }
  return bytes, nil
}

func (as* Assembler) AssembleFile(fileName string) ([]byte, error) {
  f, err := os.Open(fileName)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  bytes := make([]byte, 0)

  for _, b := range []bool{true, false} {
    scanner := bufio.NewScanner(f)
	  scanner.Split(bufio.ScanLines)
    bytes, err = as.parse(scanner, b)
  }

  fmt.Println(as.labelIndex)
  fmt.Println(bytes)

  return bytes, err
}
