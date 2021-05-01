package main

import (
	"fmt"
)

func main() {
	var (
		i   int   = 0
		i8  int8  = 8
		i16 int16 = 16
		i32 int32 = 32
		i64 int64 = 64
	)
	fmt.Println(i, i8, i16, i32, i64)

	var (
		u   uint   = 0
		u8  uint8  = 8
		u16 uint16 = 16
		u32 uint32 = 32
		u64 uint64 = 64
	)
	fmt.Println(i, i8, i16, i32, i64)
	fmt.Println(u, u8, u16, u32, u64)

}
