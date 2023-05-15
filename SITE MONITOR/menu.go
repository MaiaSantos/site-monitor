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

const monitoramentos = 3 // quantos sites serão monitorados
const delay = 5          // intervalo de tempo de um site para o outro no momento de ser monitorado

func main() {

	exibeIntro()

	for {

		exibeMenu()
		comando := insercaoInfo()
		switch comando {
		case 0:
			encerrarPrograma()
			os.Exit(0)
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		default:
			fmt.Println("Comando não encontrado")
			fmt.Println("\n")
			os.Exit(-1)
		}
	}

}

func exibeIntro() {

	nome := "João"
	versao := 1.1

	fmt.Println("Olá sr.", nome)
	fmt.Println("\n")

	fmt.Println("Este programa está na versão", versao)
	fmt.Println("\n")
}

func exibeMenu() {

	fmt.Println("0 - Sair do programa")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("\n")

	fmt.Println("Digite a opção desejada:")
	fmt.Println("\n")

}

func insercaoInfo() int {

	var comando int
	fmt.Scanf("%d", &comando) /* - Scanf serve para descobrir o tipo da variável, se colocar somente Scan nesse caso, a var terá que ser int
	- & serve para ver o endereço da var e alocar um valor*/
	fmt.Println("\n")

	fmt.Println("O comando escolhido foi o:", comando)
	fmt.Println("\n")
	//fmt.Println("O endereço da minha variável comando é", &comando)

	return comando

}

func encerrarPrograma() {
	fmt.Println("Programa encerrado")
	fmt.Println("\n")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	fmt.Println("\n")

	/* sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br",
	"https://www.caelum.com.br"} */

	sites := leSites()
	fmt.Println("\n")

	fmt.Println(sites)
	fmt.Println("\n")

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", i, ":", site)
			testaSite(site)
			fmt.Println("\n")

		}
		time.Sleep(delay * time.Second)
	}

}

func exibirLogs() {
	fmt.Println("Carregando logs")
	arquivo, err := ioutil.ReadFile("log.txt") // lendo o arquivo

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo)) // passando em string para os logs serem exibidos como texto
	fmt.Println("\n")
}

func testaSite(site string) {
	resp, err := http.Get(site) // pega a url do site
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 { // código de sucesso do site
		fmt.Println("Site:", site, ", foi carregado com sucesso!")
		registraLog(site, true)
		fmt.Println("\n")
	} else {
		fmt.Println("Site:", site, ", não foi carregado com sucesso. Status Code:", resp.StatusCode)
		registraLog(site, false)
		fmt.Println("\n")
	}
}

func leSites() []string {

	var sites []string

	/* arquivo, err := ioutil.ReadFile("site.txt") */

	arquivo, err := os.Open("site.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // Le, cria o arquivo se caso não existir e organiza o texto

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
