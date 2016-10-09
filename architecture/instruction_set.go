package architecture

/* In enumerations */
const (

	/* Arithemtic instruction */
	INS_TYPE_ARITH = iota

	/* Control flow instruction */
	INS_TYPE_CONTROL

	/* Data movement instruction */
	INS_TYPE_MOVE

	/* Logical operation */
	INS_TYPE_LOGIC
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

const (
	INS_ADD = iota
	INS_SUB
	INS_MUL

	INS_ADDI
	INS_SUBI
	INS_MULI

	INS_LEA

	INS_MOV
	INS_MOVI

	INS_LDR
	INS_STR

	INS_JEQ
	INS_JNE
	INS_JL
	INS_JGE

	INS_CMP
)

const (
  REG_A = iota
  REG_B
  REG_C
  REG_D
  REG_W
  REG_X
  REG_Y
  REG_Z
)

type InsInfo struct {
	ins_type int
	op1      int
	op2      int
	op3      int
}

type InstructionSet struct {
  ins_map map[int]InsInfo
  ins_str map[string]int
  reg_str map[string]int
}


func (is* InstructionSet) Init() {
  is.ins_map = make(map[int]InsInfo)
  is.ins_str = make(map[string]int)
  is.reg_str = make(map[string]int)


  is.insMapInit()
  is.strMapsInit()
}

func (is* InstructionSet) InsStrDecode(ins string) int {
  if val, ok := is.ins_str[ins]; ok {
    return val
  }
  return -1
}

func (is* InstructionSet) InsIdDecode(id int) InsInfo {
  if val, ok := is.ins_map[id]; ok {
    return val
  }
  return InsInfo{-1, -1, -1, -1}
}

func (is* InstructionSet) RegStrDecode(reg string) int {
  if val, ok := is.reg_str[reg]; ok {
    return val
  }
  return -1
}

func (is* InstructionSet) strMapsInit() {
  is.ins_str["add"] = INS_ADD
  is.ins_str["sub"] = INS_SUB
  is.ins_str["mul"] = INS_MUL

  is.ins_str["addi"] = INS_ADDI
  is.ins_str["subi"] = INS_SUBI
  is.ins_str["muli"] = INS_MULI

  is.ins_str["lea"] = INS_LEA

  is.ins_str["mov"] = INS_MOV
  is.ins_str["movi"] = INS_MOVI

  is.ins_str["ldr"] = INS_LDR
  is.ins_str["str"] = INS_STR

  is.ins_str["jeq"] = INS_JEQ
  is.ins_str["jne"] = INS_JNE
  is.ins_str["jl"] = INS_JL
  is.ins_str["jge"] = INS_JGE

  is.ins_str["cmp"] = INS_CMP

  /* Reg maps */

  is.reg_str["a"] = REG_A
  is.reg_str["b"] = REG_B
  is.reg_str["c"] = REG_C
  is.reg_str["d"] = REG_D
  is.reg_str["w"] = REG_W
  is.reg_str["x"] = REG_X
  is.reg_str["y"] = REG_Y
  is.reg_str["z"] = REG_Z
}

func (is* InstructionSet) insMapInit() {

	/* Arithmetic */
	is.ins_map[INS_ADD] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
	is.ins_map[INS_SUB] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
	is.ins_map[INS_MUL] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}

	is.ins_map[INS_ADDI] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}
	is.ins_map[INS_SUBI] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}
	is.ins_map[INS_MULI] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}

	is.ins_map[INS_LEA] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_ADDRESS, OP_CONSTANT}

	/* Data movement */
	is.ins_map[INS_MOV] = InsInfo{INS_TYPE_MOVE, OP_REGISTER, OP_REGISTER, OP_EMPTY}
	is.ins_map[INS_MOVI] = InsInfo{INS_TYPE_MOVE, OP_REGISTER, OP_CONSTANT, OP_EMPTY}

	is.ins_map[INS_LDR] = InsInfo{INS_TYPE_MOVE, OP_REGISTER, OP_REGISTER, OP_EMPTY}
	is.ins_map[INS_STR] = InsInfo{INS_TYPE_MOVE, OP_REGISTER, OP_REGISTER, OP_EMPTY}

	/* Control flow */

	is.ins_map[INS_JEQ] = InsInfo{INS_TYPE_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
	is.ins_map[INS_JNE] = InsInfo{INS_TYPE_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
	is.ins_map[INS_JL] = InsInfo{INS_TYPE_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
	is.ins_map[INS_JGE] = InsInfo{INS_TYPE_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}

	/* Logical */

	is.ins_map[INS_CMP] = InsInfo{INS_TYPE_LOGIC, OP_REGISTER, OP_REGISTER, OP_EMPTY}

}
