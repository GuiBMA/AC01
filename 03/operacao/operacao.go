package operacao

import (
	ctt "03/contatos"
	"fmt"
)

func AdicionaContato(c ctt.Contato, a *[5]ctt.Contato) {
	for ind, contato := range a {
		if (contato == ctt.Contato{}) {
			a[ind] = c
			break
		}
	}
}

func ExcluiContato(a *[5]ctt.Contato) {
	if (a[0] == ctt.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
	}

	for ind, contato := range a {
		if (contato == ctt.Contato{}) {
			a[ind-1] = ctt.Contato{}
		}
	}
}

func EditaEmail(c *[5]ctt.Contato, i int, novoEmail string) {
	if i < 0 || i >= len(c) {
		fmt.Println("Índice inválido")
	}
	c[i].Email = novoEmail
}
