package operacao

import (
	"fmt"
	ctt "03AC/contatos"
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

func EditaEmail(indice, novoEmail, c *[5]ctt.Contato) {
	c[indice].AlterarEmail(novoEmail)
}
