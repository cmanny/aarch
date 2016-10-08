package architecture

/* In enumerations */
const (

  /* Arithemtic instruction */
  INS_ARITH = iota

  /* Control flow instruction */
  INS_CONTROL

  /* Data movement instruction */
  INS_MOVE

  /* Logical operation */
  INS_LOGIC
)

/* Operand enumerations */
const (

  /* No instruction in this slot */
  OP_EMPTY = iota

  /* Memory address constant */
  OP_ADDRESS

  /* Register identifier */
  OP_REGISTER

  /* Constant value */
  OP_CONSTANT
)

type InsInfo struct {
  ins_type int
  op1 int
  op2 int
  op3 int
}

var ins_map = make(map[string]InsInfo)

func MappingInit() {

  /* Arithmetic */
  ins_map["add"] = InsInfo{INS_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
  ins_map["sub"] = InsInfo{INS_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
  ins_map["mul"] = InsInfo{INS_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}

  ins_map["addi"] = InsInfo{INS_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}
  ins_map["subi"] = InsInfo{INS_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}
  ins_map["muli"] = InsInfo{INS_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}

  ins_map["lea"] = InsInfo{INS_ARITH, OP_REGISTER, OP_ADDRESS, OP_CONSTANT}

  /* Data movement */
  ins_map["mov"] = InsInfo{INS_MOVE, OP_REGISTER, OP_REGISTER, OP_EMPTY}
  ins_map["movi"] = InsInfo{INS_MOVE, OP_REGISTER, OP_CONSTANT, OP_EMPTY}

  ins_map["load"] = InsInfo{INS_MOVE, OP_REGISTER, OP_ADDRESS, OP_EMPTY}
  ins_map["stor"] = InsInfo{INS_MOVE, OP_ADDRESS, OP_REGISTER, OP_EMPTY}

  /* Control flow */

  ins_map["jeq"] = InsInfo{INS_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
  ins_map["jne"] = InsInfo{INS_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
  ins_map["jl"] = InsInfo{INS_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
  ins_map["jge"] = InsInfo{INS_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}

  /* Logical */

  ins_map["cmp"] = InsInfo{INS_LOGIC, OP_REGISTER, OP_REGISTER, OP_EMPTY}

}
