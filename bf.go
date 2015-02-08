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
	fmt.Println("status would go here")
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
func (vm *bfm) ex(nm rune) {
	switch nm {
	case '+':
		vm.tape[vm.cpu.dc] += 1
	case '-':
		vm.tape[vm.cpu.dc] -= 1
	case '.':
		fmt.Print(int(vm.tape[vm.cpu.dc]))
	default:
		fmt.Println("invalid instrucion" + string(nm))
	}
	vm.cpu.pc += 1

}
func (vm *bfm) step() {
	vm.ex(rune(vm.program[vm.cpu.pc]))
}
func main() {
	vm := new(bfm)
	vm.create(1000, 1000, "+.")
	vm.step()
	vm.step()
	vm.cpu.status()

}
