.code
  movi a 0x00
  movi b 0x01
  movi c 0x0a
.loopstart
  mov  d b
  add  b b a
  mov  a a
  subi c 0x01
  jne  .loopstart
  halt

.data
