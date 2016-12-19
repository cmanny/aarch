.code
  movi a 0
  movi b 1
  movi c 40
  jmp .complex
.simple
  mov  d b
  add  b b a
  mov  a d
  mov  w b
  subi c c 1
  cmpi d c 0
  jne .simple d
  halt
.complex
  xor  a a b
  xor  b a b
  xor  a a b
  add  b a b
  mov  w b
  subi c c 1
  cmpi d c 0
  jne .complex d
  halt

.data
