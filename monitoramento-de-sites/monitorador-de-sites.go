package main

import "fmt"

func main() {

	fmt.Println("1 - Monitorar sites")
	fmt.Println("2 - Logs")
	fmt.Println("\n0 - Sair")
	fmt.Println("\nDigite a opção desejada:")

	var escolha int
	// recebe o valor digitado e armazena na variavel, recebe como parametro o ponteiro da variavel onde será armazenada, indicado com o &
	fmt.Scan(&escolha)

	//ifs sem parenteses
	if escolha == 1 {
		fmt.Println("Monitorando... ")
	} else if escolha == 2 {
		fmt.Println("Logs: ")
	} else if escolha == 0 {
		fmt.Println("Saindo...")
	} else {
		fmt.Println("Escolha inválida")
	}
}
