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

type InsInfo struct {
	ins_type int
	op1      int
	op2      int
	op3      int
}

var ins_map = make(map[int]InsInfo)

func MappingInit() {

	/* Arithmetic */
	ins_map[INS_ADD] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
	ins_map[INS_SUB] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
	ins_map[INS_MUL] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}

	ins_map[INS_ADDI] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}
	ins_map[INS_SUBI] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}
	ins_map[INS_MULI] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_CONSTANT, OP_EMPTY}

	ins_map[INS_LEA] = InsInfo{INS_TYPE_ARITH, OP_REGISTER, OP_ADDRESS, OP_CONSTANT}

	/* Data movement */
	ins_map[INS_MOV] = InsInfo{INS_TYPE_MOVE, OP_REGISTER, OP_REGISTER, OP_EMPTY}
	ins_map[INS_MOVI] = InsInfo{INS_TYPE_MOVE, OP_REGISTER, OP_CONSTANT, OP_EMPTY}

	ins_map[INS_LDR] = InsInfo{INS_TYPE_MOVE, OP_REGISTER, OP_REGISTER, OP_EMPTY}
	ins_map[INS_STR] = InsInfo{INS_TYPE_MOVE, OP_REGISTER, OP_REGISTER, OP_EMPTY}

	/* Control flow */

	ins_map[INS_JEQ] = InsInfo{INS_TYPE_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
	ins_map[INS_JNE] = InsInfo{INS_TYPE_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
	ins_map[INS_JL] = InsInfo{INS_TYPE_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}
	ins_map[INS_JGE] = InsInfo{INS_TYPE_CONTROL, OP_ADDRESS, OP_EMPTY, OP_EMPTY}

	/* Logical */

	ins_map[INS_CMP] = InsInfo{INS_TYPE_LOGIC, OP_REGISTER, OP_REGISTER, OP_EMPTY}

}
