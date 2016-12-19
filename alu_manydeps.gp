.code
  movi z 40
.loop
  addi a a 2
  addi b a 3
  addi c a 4
  addi d a 5
  addi x b 2
  addi y b 3
  mul a a c
  mul b d b
  mul c c x
  addi a a 2
  addi b a 3
  addi a a 2
  addi b a 3
  addi c a 4
  addi d a 5
  addi x b 2
  addi y b 3
  mul a a c
  mul b d b
  mul c c x
  addi a a 2
  addi b a 3
  addi a a 2
  addi b a 3
  addi c a 4
  addi d a 5
  addi x b 2
  addi y b 3
  mul a a c
  mul b d b
  mul c c x
  addi a a 2
  addi b a 3
  addi a a 2
  addi b a 3
  addi c a 4
  addi d a 5
  addi x b 2
  addi y b 3
  mul a a c
  mul b d b
  mul c c x
  addi a a 2
  addi b a 3
  mul a a a
  mul b b b
  mul c c c
  mul d d d
  mul b b b
  mul c c c
  subi z z 1
  cmpi w z 0
  jne .loop w
  halt



.data
