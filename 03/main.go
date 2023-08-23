package main

import (
	ct "03/contatos"
	op "03/operacao"
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
		fmt.Print("Digite (1)adicionar, (2)remover, (3)exibir, (4)editar caso queira sair qualquer tecla: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			op.AdicionaContato(ct.Contato{Nome: nome, Email: email}, &contatos)

		case "2":
			op.ExcluiContato(&contatos)

		case "3":
			fmt.Println(contatos)
			for i, contato := range contatos {
				fmt.Print(i, " Nome: ", contato.Nome, i, " Email: ", contato.Email, "\n")
			}

		case "4":
			fmt.Print("Digite o índice do contato que deseja editar: ")
			var i int
			fmt.Scanln(&i)

			if i >= 0 && i < len(contatos) {
				fmt.Print("Digite o novo e-mail: ")
				var novoEmail string
				fmt.Scanln(&novoEmail)

				op.EditaEmail(&contatos, i, novoEmail)
			} else {
				fmt.Println("Índice inválido!")
			}

		default:
			return
		}

		fmt.Println(contatos)
	}
}
