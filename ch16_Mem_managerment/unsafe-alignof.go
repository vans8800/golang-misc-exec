package main

import (
	"fmt"
	"unsafe"
)

type Country struct {
	isAdvanced bool
	id         int16
	population int64
}

func main() {

	var b byte
	var f float32

	fmt.Println("byte类型的对齐边界:", unsafe.Alignof(b))
	fmt.Println("byte指针类型的对齐边界:", unsafe.Alignof(&b))
	fmt.Println("float32类型的对齐边界:", unsafe.Alignof(f))
	fmt.Println("float32指针类型的对齐边界:", unsafe.Alignof(&f))

	fmt.Println("\n")

	var country = Country{}

	fmt.Println("对齐边界", unsafe.Alignof(country))
	fmt.Println("占用空间", unsafe.Sizeof(country))
}
