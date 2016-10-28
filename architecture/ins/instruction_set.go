package ins

import (
	"fmt"
)

/* In enumerations */
const (

	/* Arithemtic instruction */
	TYPE_ARITH = iota

	/* Control flow instruction */
	TYPE_CONTROL

	/* Data movement instruction */
	TYPE_MOVE

	/* Logical operation */
	TYPE_LOGIC

	TYPE_IO
)

/* Operand enumerations */
const (

	/* No instruction in this slot */
	OP_EMPTY = iota

	/* Memory address constant */
	OP_ADDRESS_8

	/* Register identifier */
	OP_REGISTER

	/* Indirect address identifier */
	OP_IND_ADDR

	/* Constant value */
	OP_CONSTANT_8
)

const (
	HALT = iota

	ADD
	SUB
	MUL
	XOR

	ADDI
	SUBI
	MULI
	XORI

	LEAL
	LEAH

	MOV
	MOVI

	LDR
	STR

	/** Use local jumps and absolute jumps, so that the processor can
	    make better guesses at which data to load into the cache and pipeline
	**/

	JMP
	JEQ
	JNE
	JL
	JGE
	JRND

	AJMP
	AJEQ
	AJNE
	AJL
	AJGE

	CMP
	CMPI

	IOIO
	IOII
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
	Ins_type int
	Op1      int
	Op2      int
	Op3      int
}

type InstructionSet struct {
	ins_map map[byte]InsInfo
	ins_str map[string]byte
	reg_str map[string]byte
}

func (is *InstructionSet) Init() {
	is.ins_map = make(map[byte]InsInfo)
	is.ins_str = make(map[string]byte)
	is.reg_str = make(map[string]byte)

	is.insMapInit()
	is.strMapsInit()
}

func (is *InstructionSet) InsStrDecode(ins string) (byte, error) {
	if val, ok := is.ins_str[ins]; ok {
		return val, nil
	}
	return 255, fmt.Errorf("could not find instruction string in mapping (%s)", ins)
}

func (is *InstructionSet) InsIdDecode(id byte) (InsInfo, error) {
	if val, ok := is.ins_map[id]; ok {
		return val, nil
	}
	return InsInfo{-1, -1, -1, -1}, fmt.Errorf("could not find instruction id in mapping (%d)", id)
}

func (is *InstructionSet) RegStrDecode(reg string) (byte, error) {
	if val, ok := is.reg_str[reg]; ok {
		return val, nil
	}
	return 255, fmt.Errorf("could not find reg str in mapping (%s)", reg)
}

func (is *InstructionSet) strMapsInit() {

	is.ins_str["add"] = ADD
	is.ins_str["sub"] = SUB
	is.ins_str["mul"] = MUL
	is.ins_str["xor"] = XOR

	is.ins_str["addi"] = ADDI
	is.ins_str["subi"] = SUBI
	is.ins_str["muli"] = MULI
	is.ins_str["xori"] = XORI

	is.ins_str["leal"] = LEAL
	is.ins_str["leah"] = LEAH

	is.ins_str["mov"] = MOV
	is.ins_str["movi"] = MOVI

	is.ins_str["ldr"] = LDR
	is.ins_str["str"] = STR

	is.ins_str["jmp"] = JMP
	is.ins_str["jeq"] = JEQ
	is.ins_str["jne"] = JNE
	is.ins_str["jl"] = JL
	is.ins_str["jge"] = JGE
	is.ins_str["jrnd"] = JRND

	is.ins_str["ajmp"] = AJMP
	is.ins_str["ajeq"] = AJEQ
	is.ins_str["ajne"] = AJNE
	is.ins_str["ajl"] = AJL
	is.ins_str["ajge"] = AJGE

	is.ins_str["cmp"] = CMP
	is.ins_str["cmpi"] = CMPI
	is.ins_str["halt"] = HALT

	is.ins_str["ioio"] = IOIO
	is.ins_str["ioii"] = IOII

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

func (is *InstructionSet) insMapInit() {

	/* Arithmetic */
	is.ins_map[ADD] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
	is.ins_map[SUB] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
	is.ins_map[MUL] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}
	is.ins_map[XOR] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_REGISTER}

	is.ins_map[ADDI] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_CONSTANT_8}
	is.ins_map[SUBI] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_CONSTANT_8}
	is.ins_map[MULI] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_CONSTANT_8}
	is.ins_map[XORI] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_REGISTER, OP_CONSTANT_8}

	is.ins_map[LEAL] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_CONSTANT_8, OP_CONSTANT_8}
	is.ins_map[LEAH] = InsInfo{TYPE_ARITH, OP_REGISTER, OP_CONSTANT_8, OP_CONSTANT_8}

	/* Data movement */
	is.ins_map[MOV] = InsInfo{TYPE_MOVE, OP_REGISTER, OP_REGISTER, OP_EMPTY}
	is.ins_map[MOVI] = InsInfo{TYPE_MOVE, OP_REGISTER, OP_CONSTANT_8, OP_EMPTY}

	is.ins_map[LDR] = InsInfo{TYPE_MOVE, OP_REGISTER, OP_IND_ADDR, OP_EMPTY}
	is.ins_map[STR] = InsInfo{TYPE_MOVE, OP_IND_ADDR, OP_REGISTER, OP_EMPTY}

	/* Control flow */

	is.ins_map[JMP] = InsInfo{TYPE_CONTROL, OP_ADDRESS_8, OP_EMPTY, OP_EMPTY}
	is.ins_map[JEQ] = InsInfo{TYPE_CONTROL, OP_ADDRESS_8, OP_REGISTER, OP_EMPTY}
	is.ins_map[JNE] = InsInfo{TYPE_CONTROL, OP_ADDRESS_8, OP_REGISTER, OP_EMPTY}
	is.ins_map[JL] = InsInfo{TYPE_CONTROL, OP_ADDRESS_8, OP_REGISTER, OP_EMPTY}
	is.ins_map[JGE] = InsInfo{TYPE_CONTROL, OP_ADDRESS_8, OP_REGISTER, OP_EMPTY}
	is.ins_map[JRND] = InsInfo{TYPE_CONTROL, OP_ADDRESS_8, OP_REGISTER, OP_EMPTY}

	is.ins_map[AJMP] = InsInfo{TYPE_CONTROL, OP_IND_ADDR, OP_EMPTY, OP_EMPTY}
	is.ins_map[AJEQ] = InsInfo{TYPE_CONTROL, OP_IND_ADDR, OP_REGISTER, OP_EMPTY}
	is.ins_map[AJNE] = InsInfo{TYPE_CONTROL, OP_IND_ADDR, OP_REGISTER, OP_EMPTY}
	is.ins_map[AJL] = InsInfo{TYPE_CONTROL, OP_IND_ADDR, OP_REGISTER, OP_EMPTY}
	is.ins_map[AJGE] = InsInfo{TYPE_CONTROL, OP_IND_ADDR, OP_REGISTER, OP_EMPTY}

	is.ins_map[HALT] = InsInfo{TYPE_CONTROL, OP_EMPTY, OP_EMPTY, OP_EMPTY}

	/* Logical */

	is.ins_map[CMP] = InsInfo{TYPE_LOGIC, OP_REGISTER, OP_REGISTER, OP_REGISTER}

	is.ins_map[CMPI] = InsInfo{TYPE_LOGIC, OP_REGISTER, OP_REGISTER, OP_CONSTANT_8}

	/* IO */

	is.ins_map[IOIO] = InsInfo{TYPE_IO, OP_REGISTER, OP_REGISTER, OP_CONSTANT_8}
	is.ins_map[IOII] = InsInfo{TYPE_IO, OP_REGISTER, OP_REGISTER, OP_CONSTANT_8}
}
