package main

import "github.com/s2607/mnp"

//import "io/ioutil"
//import "strconv"
func main() {
	vm := new(mnp.Task)
	vm.Create(1000, 1000, "+++++.-.++++++.[-.]+.")
	for !vm.Step() {
		vm.Status()
	}
	//vm.cpu.Status()

}
