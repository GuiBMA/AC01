package main

import (
	ct "03AC/contatos"
	op "03AC/operacao"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var contatos [5]ct.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) adicionar, (2) remover (3) exibir ou qualquer tecla para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			contatos = op.AdicionaContato(ct.Contato{Nome: nome, Email: email}, &contatos)

		case "2":
			contatos = op.ExcluiContato(&contatos)

		case "3":
			fmt.Println(contatos)

		default:
			return
		}

		fmt.Println(contatos)
	}
}
