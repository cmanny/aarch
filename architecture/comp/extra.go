package comp

type InsIn struct {

  Ip int
  Code byte
  RawOp1 byte
  RawOp2 byte
  RawOp3 byte

  Op1 int
  Op2  int
  Op3  int
  Valid bool
  Speculative bool

  Result  int

}
