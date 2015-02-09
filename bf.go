package main

import "fmt"

//import "io/ioutil"
//import "strconv"

type alu struct {
	pc     int
	dc     int
	lc     int
	halted bool
}

func (cpu alu) status() {
	fmt.Printf("%d\t  %d\t  %d\n", cpu.dc, cpu.pc, cpu.lc)
}

type bfm struct {
	cpu     alu
	stack   []int
	tape    []int
	program string
}

func (vm *bfm) create(ssize int, tsize int, program string) {
	vm.stack = make([]int, ssize)
	vm.tape = make([]int, tsize)
	vm.program = program
	vm.cpu.dc = 0
	vm.cpu.lc = 0
	vm.cpu.pc = 0
	vm.cpu.halted = false
}

func (vm *bfm) skip() {
	for vm.program[vm.cpu.pc] != ']' {
		vm.cpu.pc += 1
		if vm.program[vm.cpu.pc] == '[' {
			vm.skip()
		}
	}
	vm.cpu.lc -= 1
	//vm.cpu.pc += 2
}
func (vm *bfm) ex(nm rune) {
	switch nm {
	case '+':
		vm.tape[vm.cpu.dc] += 1
	case '-':
		vm.tape[vm.cpu.dc] -= 1
	case '.':
		//fmt.Print(int(vm.tape[vm.cpu.dc]))
	case '<':
		vm.cpu.dc -= 1
	case '>':
		vm.cpu.dc += 1
	case '[':
		vm.cpu.lc += 1
		vm.stack[vm.cpu.lc] = vm.cpu.pc - 1
		if vm.tape[vm.cpu.dc] == 0 {
			vm.skip()
		}
	case ']':
		vm.cpu.pc = vm.stack[vm.cpu.lc]
		vm.cpu.lc -= 1

	}
	vm.cpu.pc += 1
}
func (vm *bfm) step() bool {
	vm.cpu.halted = (vm.cpu.pc >= len(vm.program)-1)
	if !vm.cpu.halted {
		vm.ex(rune(vm.program[vm.cpu.pc])) //TODO:UTF8
	}
	return vm.cpu.halted
}
func (vm bfm) status() {
	if !vm.cpu.halted {
		fmt.Printf("%d\t %d\t %c\t ", vm.stack[vm.cpu.lc], vm.tape[vm.cpu.dc], vm.program[vm.cpu.pc])
		vm.cpu.status()
	}
}
func main() {
	vm := new(bfm)
	vm.create(1000, 1000, "+++++.-.++++++.[-.]+.")
	for !vm.step() {
		vm.status()
	}
	vm.cpu.status()

}
