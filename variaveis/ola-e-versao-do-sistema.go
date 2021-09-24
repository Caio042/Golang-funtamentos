package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*Inferencia de tipo, declaração e atribuição de variavel com := */
	nome := "Caio"
	var versao float32 = 0.9
	fmt.Println("Olá,", nome)
	fmt.Println("Você está utilizando o programa na versão", versao)
	fmt.Println("A variável nome é do tipo", reflect.TypeOf(nome))
}
