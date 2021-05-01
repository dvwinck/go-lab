package main

import "fmt"

func main() {
	var variavel1 string = "Variavel1 tipo string"
	variavel2 := "variavel implicita"

	var (
		inteiro int     = 20
		real    float32 = 1.0
		texto   string  = "teste"
	)
	var1, var2 := "1", "2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(inteiro)
	fmt.Println(real)
	fmt.Println(texto)
	fmt.Println(var1, var2)

	var1, var2 = var2, var1
	fmt.Println(var1, var2)

}
