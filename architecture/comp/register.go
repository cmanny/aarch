package comp

import (
	"bytes"
	"fmt"
)

type RegisterFile struct {
	Communicator
	regs  [32]int
	Flags int
}

func (rf *RegisterFile) Init() {
	rf.InitComms()

	rf.Inputs["in1"] = make(chan interface{}, 1)
	rf.Outputs["out1"] = make(chan interface{}, 1)
}

func (rf *RegisterFile) Data() interface{} {
	return ""
}

func (rf *RegisterFile) State() string {
	return ""
}

func (rf *RegisterFile) Cycle() {
	rf.Outputs["out1"] <- rf.regs[(<-rf.Inputs["in1"]).(int)]
}

func (rf *RegisterFile) Contents() {
	var buffer bytes.Buffer
	for i := 0; i < 32; i++ {
		buffer.WriteString("r" + fmt.Sprintf("%d", i) + ": " + fmt.Sprintf("%d", rf.regs[i]) + "\n")
	}
	fmt.Println(buffer.String())
}
