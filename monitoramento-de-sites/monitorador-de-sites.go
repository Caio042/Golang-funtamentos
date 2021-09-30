package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
		case 2:
			imprimeLog()
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
	sites := leSitesDoArquivo()

	for i := 0; i < quantidadeMonitoramento; i++ {
		time.Sleep(delay * time.Second)
		// a função retorna dois valores, mas só o primeiro me interessa
		// ignoro o segundo com _. Também da pra ignorar o primeiro com o mesmo simbolo
		for _, site := range sites {
			testarSite(site)
		}
		fmt.Println("\n------------")
	}
}

func testarSite(site string) {
	resposta, erro := http.Get(site)

	if erro != nil {
		fmt.Println("Deu erro: ", erro)
	}

	if resposta.StatusCode == 200 {
		fmt.Println(site, "- Está funcionando, teoricamente")
	} else {
		fmt.Println(site, "- Caiu bixo, Status:", resposta.StatusCode)
	}

	escreveLog(site, resposta.StatusCode)
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, erro := os.Open("resources/sites.txt")

	if erro != nil {
		fmt.Println("Erro ao ler arquivo:", erro)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if erro == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func escreveLog(site string, status int) {
	arquivo, erro := os.OpenFile("resources/log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println("Erro ao abrir log", erro)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") +
		" --- " + site + " --- STATUS: " + strconv.Itoa(status) + "\n")

	arquivo.Close()
}

func imprimeLog() {
	arquivo, _ := ioutil.ReadFile("resources/log.txt")

	fmt.Println(string(arquivo))
}

func sair() {
	os.Exit(0)
}
