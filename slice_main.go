package main

import (
	"fmt"

	"github.com/MinhAndy/ica02/is105-ica02/slice"
)

func main() {
	var byteSlice1 = slice.AllocateMake(100)
	fmt.Println("&byteslice1[0]")
	fmt.Println(&byteSlice1[0])
	aslice := slice.Reslice(byteSlice1, 50, 100)
	fmt.Println("&aslice[0]")
	fmt.Println(&aslice[0])
	fmt.Println("&byteslice1[50]")
	fmt.Println(&byteSlice1[50])
}
