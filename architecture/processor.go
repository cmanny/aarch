package architecture

import (
	"fmt"

	"github.com/cmanny/aarch/architecture/comp"
	"github.com/cmanny/aarch/architecture/comp/exe"
	"github.com/cmanny/aarch/architecture/ins"
	//"bufio"
)

type Processor struct {
	clockSpeed int
	numExUnits int
	printDebug bool
	exit       bool

	ip int

	is *ins.InstructionSet

	/* Components */

	mem *comp.Memory

	fu *comp.NullComponent
	du *comp.NullComponent

	cu *exe.ControlUnit
	au *exe.ArithmeticUnit
	lu *exe.LogicUnit

	rf *comp.RegisterFile
}

/**
  Private methods of Processor
**/

func (p *Processor) preRun() {
	if p.printDebug {
		fmt.Println("Debug ON")
	}
}

func (p *Processor) fetch() {
	p.mem.In(p, "in1") <- p.ip
	bytes := <-p.mem.Out(p, "out1")
	fmt.Println(bytes)

	p.cu.In(p.mem, "ins") <- bytes
	p.ip = (<-p.cu.Out(p, "ip")).(int)
	if p.ip == -1 {
		// bail
		p.exit = true
	}

}

func (p *Processor) decode() {

}

func (p *Processor) execute() {
}

func (p *Processor) writeback() {

}

/**
  Public methods of Processor
**/

func (p *Processor) Init(is *ins.InstructionSet, mem *comp.Memory) {
	comp.Init()
	p.is = is

	/* Init all sub components */
	p.mem = mem

	p.fu = &comp.NullComponent{}
	p.du = &comp.NullComponent{}

	p.cu = &exe.ControlUnit{}
	p.cu.Init()

	p.au = &exe.ArithmeticUnit{}
	p.au.Init()

	p.lu = &exe.LogicUnit{}
	p.lu.Init()

	p.ip = (<-p.cu.Out(p, "ip")).(int)

	comp.AddAll(
		&comp.CompWrapper{
			Name: "RAM",
			Obj:  p.mem,
		},
		&comp.CompWrapper{
			Name: "Fetch",
			Obj:  p.fu,
		},
		&comp.CompWrapper{
			Name: "Decode",
			Obj:  p.du,
		},
		&comp.CompWrapper{
			Name: "ControlUnit",
			Obj:  p.cu,
		},
		&comp.CompWrapper{
			Name: "ArithmeticUnit",
			Obj:  p.au,
		},
		&comp.CompWrapper{
			Name: "LogicUnit",
			Obj:  p.lu,
		},
	)
}

func (p *Processor) Data() interface{} {
	return 1
}

func (p *Processor) State() string {
	return ""
}

func (p *Processor) Cycle() {

}

func (p *Processor) Debug(toggle bool) {
	p.printDebug = toggle
}

func (p *Processor) SetIP(ip int) {
	p.ip = ip
}

func (p *Processor) Run() {
	p.preRun()
	fmt.Println("Processor beginning")

	for {
		for _, c := range comp.Comps {
			go c.Obj.Cycle()
		}
		p.fetch()
		p.decode()
		p.execute()
		p.writeback()

		if p.exit {
			return
		}
	}
}
