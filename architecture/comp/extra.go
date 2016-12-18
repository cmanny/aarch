package comp

type InsIn struct {

  Tag int
  Ip int
  Code byte
  RawOp1 byte
  RawOp2 byte
  RawOp3 byte

  Op1 int
  Op2  int
  Op3  int

  Op1Tag int
  Op2Tag int
  Op3Tag int

  Op1Valid bool
  Op2Valid bool
  Op3Valid bool

  Valid bool
  FlowId int

  Result  int

}
