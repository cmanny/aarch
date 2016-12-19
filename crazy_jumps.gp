.code
  movi a 0
  movi b 1
  movi c 10
  jmp .loop1
.secret
  addi d d 1
  cmpi w d 20
  jeq .end w
  muli y y 1
  jmp .mad
.loop1
  addi a a 1
  addi b b 1
  subi c c 1
  cmpi z c 0
  jne .loop1 z
  jl .secret z
.mad
  jmp .secret
.end
  halt

.data
