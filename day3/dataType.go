package main

import (
	"errors"
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

	var (
		u   uint   = 0
		u8  uint8  = 8
		u16 uint16 = 16
		u32 uint32 = 32
		u64 uint64 = 64
	)
	fmt.Println(i, i8, i16, i32, i64)
	fmt.Println(u, u8, u16, u32, u64)

	var exemploRune rune = 25
	var exemploByte byte = 25

	fmt.Println(exemploRune, exemploByte)

	var f32 float32 = 3.14
	var f64 float64 = 3.1415

	fmt.Println(f32, f64)

	var str1 string = "value 1 "
	str2 := "value 2 "
	char := 'D'
	print(str1, str2, char, " ")

	var boleano bool = true
	fmt.Println(boleano)

	var erroExemplo error = errors.New("erro interno")
	fmt.Println(erroExemplo)

}
