package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const quantidadeMonitoramento = 3
const delay = 5

func main() {

	// ta saindo dois retornos da função caraa
	texto, numero := devolveStringEInt()
	fmt.Println(texto, numero)

	// não existe while, loop infinito é assim
	for {
		exibeMenu()

		escolha := leComando()
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

		switch escolha {
		case 0:
			sair()
		case 1:
			monitorarSites()
		default:
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Monitorar sites")
	fmt.Println("2 - Logs")
	fmt.Println("\n0 - Sair")
	fmt.Println("\nDigite a opção desejada:")
}

// a função tem varios retornos caraaa
func devolveStringEInt() (string, int) {
	inteiro := 10
	cadeia := "texto"

	return cadeia, inteiro
}

func leComando() int {
	var escolha int
	// recebe o valor digitado e armazena na variavel, recebe como parametro o ponteiro da variavel onde será armazenada, indicado com o &
	fmt.Scan(&escolha)
	return escolha
}

func monitorarSites() {
	sites := []string{"http://github.com/", "http://youtube.com/", "http://twitter.com/"}

	for i := 0; i < quantidadeMonitoramento; i++ {
		time.Sleep(delay * time.Second)
		for _, site := range sites {
			testarSite(site)
		}
		fmt.Println("\n------------")
	}
	// a função retorna dois valores, mas só o primeiro me interessa
	// ignoro o segundo com _. Também da pra ignorar o primeiro com o mesmo simbolo

}

func testarSite(site string) {
	resposta, _ := http.Get(site)
	if resposta.StatusCode == 200 {
		fmt.Println(site, "- Está funcionando, teoricamente")
	} else {
		fmt.Println(site, "- Caiu bixo, Status:", resposta.StatusCode)
	}
}

func sair() {
	os.Exit(0)
}
