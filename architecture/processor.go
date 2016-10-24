package architecture

import (
	"fmt"

	"github.com/cmanny/aarch/architecture/comp"
	"github.com/cmanny/aarch/architecture/comp/exe"
	"github.com/cmanny/aarch/architecture/ins"
	//"bufio"
	"os"
)

type Processor struct {
	clockSpeed int
	numExUnits int
	printDebug bool

	ip int

	is  *ins.InstructionSet
	mem *comp.Memory
	cu  *exe.ControlUnit
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
		os.Exit(0)
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
	p.mem = mem

	p.cu = &exe.ControlUnit{}
	p.cu.Init()
	p.ip = (<-p.cu.Out(p, "ip")).(int)

	comp.AddAll(p.mem, p.cu)
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
			go c.Cycle()
		}
		p.fetch()
		p.decode()
		p.execute()
	}
}
