package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Print("teste de pacotes locais e importados")
	auxiliar.Escrever1()
	auxiliar.Escrever2()
	erro := checkmail.ValidateFormat("diogo.winck@gmail.com")
	fmt.Println(erro)

	erro = checkmail.ValidateFormat("formato invalido")
	fmt.Println(erro)

}
