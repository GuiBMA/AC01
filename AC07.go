package main

import "fmt"

type No struct {
    Key   int
    Esq  *No
    Dir *No
}

type ArvoreBinaria struct {
    Raiz *No
}

func NovaArvoreBinaria() *ArvoreBinaria {
    return &ArvoreBinaria{}
}

func (Ab *ArvoreBinaria) Insere(key int) {
    if Ab.Raiz == nil {
        Ab.Raiz = &No{Key: key}
    } else {
        Ab.Raiz = incereNo(Ab.Raiz, key)
    }
}

func incereNo(no *No, key int) *No {
    if no == nil {
        return &No{Key: key}
    }

    if key == no.Key {
        fmt.Println("Inserção inválida")
    } else if key < no.Key {
        no.Esq = incereNo(no.Esq, key)
    } else {
        no.Dir = incereNo(no.Dir, key)
    }

    return no
}

func (t *ArvoreBinaria) BuscaValor(key int) bool {
    return procuraNo(t.Raiz, key)
}

func procuraNo(no *No, key int) bool {
    if no == nil {
        return false
    }

    if key == no.Key {
        return true
    } else if key < no.Key {
        return procuraNo(no.Esq, key)
    } else {
        return procuraNo(no.Dir, key)
    }
}

func main() {
    arvore := NovaArvoreBinaria()

    arvore.Insere(5)
    arvore.Insere(3)
    arvore.Insere(7)
    arvore.Insere(2)
    arvore.Insere(4)
    arvore.Insere(6)
    arvore.Insere(8)

	fmt.Println(arvore.Raiz)
	fmt.Println(arvore.Raiz.Dir)
	fmt.Println(arvore.Raiz.Esq)
	fmt.Println(arvore.Raiz.Dir.Esq)
	fmt.Println(arvore.Raiz.Dir.Dir)
	fmt.Println(arvore.Raiz.Esq.Dir)
	fmt.Println(arvore.Raiz.Esq.Esq)
	fmt.Println("--------------------------")
    fmt.Println(arvore.BuscaValor(3))
    fmt.Println(arvore.BuscaValor(9))
}
