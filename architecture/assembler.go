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

func intStrToBytes(str string) ([]byte, error){

  val, err := strconv.Atoi(str)
  if err != nil {
    return nil, err
  }
  bs := make([]byte, 4)
  binary.LittleEndian.PutUint32(bs, uint32(val))
  return bs, nil
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

      /* if code section, continue, if data section, parse hex data */
    } else {
      /* normal instruction OR data*/
      id, err := as.is.InsStrDecode(words[0])
      if err != nil {
        return nil, err
      }
      info, err := as.is.InsIdDecode(id)
      if !labelOnly {
        bytes = append(bytes, id)
      }
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
        op_str := ""
        if len(words) >= i + 2 {
          op_str = words[i + 1]
        }
        /* Second pass we generate the bytes and add them to the return buffer */
        switch op {
          case  OP_REGISTER:
            regId, err := as.is.RegStrDecode(op_str)

            if err != nil {
              return nil, err
            }
            bytes = append(bytes, regId)
            byteIndex += 1
          case OP_CONSTANT_8:
            /* load constant into bytes */
            bs, err := intStrToBytes(op_str)
            if err != nil {
              return nil, err
            }
            bytes = append(bytes, bs[0])

            byteIndex += 1
          case OP_CONSTANT_32:
            bs, err := intStrToBytes(op_str)
            if err != nil {
              return nil, err
            }
            bytes = append(bytes, bs...)
            byteIndex += 4

          case  OP_ADDRESS_8:
            /* can be label or address, usually label */
            if op_str[0] == '.' {
              dist := as.labelIndex[op_str[1:]] - byteIndex
              bytes = append(bytes, byte(int8(dist)))
              continue
            }
            bs, err := intStrToBytes(op_str)
            if err != nil {
              return nil, err
            }
            bytes = append(bytes, bs[0])
            byteIndex += 1
            /* load address into bytes */
          case OP_ADDRESS_32:
            if op_str[0] == '.' {
              bs := make([]byte, 4)
              binary.LittleEndian.PutUint32(bs, uint32(as.labelIndex[op_str[1:]]))
              if err != nil {
                return nil, err
              }
              bytes = append(bytes,  bs...)
              continue
            }
            bs, err := intStrToBytes(op_str)
            if err != nil {
              return nil, err
            }
            bytes = append(bytes, bs...)
            byteIndex += 4
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
    f.Seek(0,0)
    scanner := bufio.NewScanner(f)
	  scanner.Split(bufio.ScanLines)
    bytes, err = as.parse(scanner, b)

    if err != nil {
      return nil, err
    }
  }

  fmt.Println(as.labelIndex)
  fmt.Println(bytes)

  return bytes, err
}
