.code
  movi a 0
  movi b 1
  movi c 10
.loopstart
  mov  d b
  add  b b a
  mov  a a
  subi c 1
  jne  .loopstart c
  halt

.data
