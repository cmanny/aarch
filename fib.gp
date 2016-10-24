.code
  movi a 0
  movi b 1
  movi c 10
  jmp .complex
.simple
  mov  d b
  add  b b a
  mov  a d
  subi c c 1
  cmpi d c 0
  jne .simple d
  halt
.complex
  xor  a a b
  xor  b a b
  xor  a a b
  add  b a b
  subi c c 1
  cmpi d c 0
  jne .complex d

.data
