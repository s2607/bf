package main

import "fmt"

//import "io/ioutil"
//import "strconv"

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

type task struct {
	cpu       alu
	callstack []int
	acumstack []int
	ldata     []int
	program   string
}
type ophandle func(*task)

func nop(vm *task) {
	vm.cpu.pc += 1
}
func (cpu *alu) initopmap() {
	cpu.opmap = make([]ophandle, int('~'))
	var i = 0
	for i = 0; i < len(cpu.opmap); i++ {
		cpu.opmap[i] = nop
	}
}

type otask struct {
	id    int
	name  string
	tasks []int //list of task ids sorted by priority

}
type system struct {
	ostate otask
	states []otask
}

func (vm *task) create(ssize int, tsize int, program string) {
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
func (vm *task) ex(nm rune) {
	var op = int(nm)
	if op < len(vm.cpu.opmap) {
		vm.cpu.opmap[op](vm)
	}
}
func (vm *task) step() bool {
	vm.cpu.halted = (vm.cpu.pc >= len(vm.program)-1)
	if !vm.cpu.halted {
		vm.ex(rune(vm.program[vm.cpu.pc])) //TODO:UTF8
	}
	return vm.cpu.halted
}
func (vm task) status() {
	if !vm.cpu.halted {
		fmt.Printf("%d\t %d\t %c\t ", vm.callstack[vm.cpu.lc], vm.ldata[vm.cpu.pc], vm.program[vm.cpu.pc])
		vm.cpu.status()
	}
}
func main() {
	vm := new(task)
	vm.create(1000, 1000, "+++++.-.++++++.[-.]+.")
	for !vm.step() {
		vm.status()
	}
	vm.cpu.status()

}
