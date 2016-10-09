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
  ins_id map[string]int
}


func (is* InstructionSet) Init() {
  is.ins_map = make(map[int]InsInfo)
  is.ins_id = make(map[string]int)


  is.insMapInit()
  is.strMapsInit()
}

func (is* InstructionSet) strMapsInit() {
  is.ins_id["add"] = INS_ADD
  is.ins_id["sub"] = INS_SUB
  is.ins_id["mul"] = INS_MUL

  is.ins_id["addi"] = INS_ADDI
  is.ins_id["subi"] = INS_SUBI
  is.ins_id["muli"] = INS_MULI

  is.ins_id["lea"] = INS_LEA

  is.ins_id["mov"] = INS_MOV
  is.ins_id["movi"] = INS_MOVI

  is.ins_id["ldr"] = INS_LDR
  is.ins_id["str"] = INS_STR

  is.ins_id["jeq"] = INS_JEQ
  is.ins_id["jne"] = INS_JNE
  is.ins_id["jl"] = INS_JL
  is.ins_id["jge"] = INS_JGE

  is.ins_id["cmp"] = INS_CMP
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
