package mnp

import "fmt"

type alu struct {
	pc     int
	lc     int
	ac     int
	halted bool
	opmap  []ophandle
}

func (cpu alu) status() {
	//fmt.Printf("%d\t  %d\t  %d\n", cpu.dc, cpu.pc, cpu.lc)
}

type Task struct {
	cpu       alu
	callstack []int
	acumstack []int
	ldata     []int
	program   string
}

type oTask struct {
	id    int
	name  string
	Tasks []int //list of Task ids sorted by priority

}
type system struct {
	ostate oTask
	states []oTask
}

func (vm *Task) Create(ssize int, tsize int, program string) {
	vm.callstack = make([]int, ssize)
	vm.acumstack = make([]int, ssize)
	vm.ldata = make([]int, tsize)
	vm.program = program
	vm.cpu.initopmap()
	vm.cpu.lc = 0
	vm.cpu.ac = 0
	vm.cpu.pc = 0
	vm.cpu.halted = false
}
func (vm *Task) ex(nm rune) {
	var op = int(nm)
	if op < len(vm.cpu.opmap) {
		vm.cpu.opmap[op](vm)
	}
}
func (vm *Task) Step() bool {
	vm.cpu.halted = (vm.cpu.pc >= len(vm.program)-1)
	if !vm.cpu.halted {
		vm.ex(rune(vm.program[vm.cpu.pc])) //TODO:UTF8
	}
	return vm.cpu.halted
}
func (vm Task) Status() {
	if !vm.cpu.halted {
		fmt.Printf("%d\t %d\t %c\t ", vm.callstack[vm.cpu.lc], vm.ldata[vm.cpu.pc], vm.program[vm.cpu.pc])
		vm.cpu.status()
	}
}
