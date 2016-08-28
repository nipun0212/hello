package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (i IPAddr) String() string {
	return fmt.Sprintf("%v", i.String())
}

func main() {
	s := IPAddr{2, 3, 5, 7}
	//i := IPAddr()
	fmt.Println(s)
}
