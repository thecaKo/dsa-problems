package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func NovaArvore() *Arvore {
	return &Arvore{}
}

func NovoNo(v int) *No {
	return &No{Valor: v}
}

func (a *Arvore) Inserir(v int) {
	a.Raiz = inserir(a.Raiz, v)
}

func inserir(no *No, v int) *No {
	if no == nil {
		return NovoNo(v)
	}
	if v < no.Valor {
		no.Esquerdo = inserir(no.Esquerdo, v)
	} else if v > no.Valor {
		no.Direito = inserir(no.Direito, v)
	}
	return no
}

func (a *Arvore) Buscar(v int) *No {
	return buscar(a.Raiz, v)
}

func buscar(no *No, v int) *No {
	if no == nil || no.Valor == v {
		return no
	}
	if v < no.Valor {
		return buscar(no.Esquerdo, v)
	}
	return buscar(no.Direito, v)
}

func (a *Arvore) BuscarIter(v int) *No {
	atual := a.Raiz
	for atual != nil {
		if v == atual.Valor {
			return atual
		}
		if v < atual.Valor {
			atual = atual.Esquerdo
		} else {
			atual = atual.Direito
		}
	}
	return nil
}

func relatar(metodo string, v int, encontrado *No) {
	if encontrado == nil {
		fmt.Printf("  %-12s valor %d: NAO encontrado na arvore\n", metodo, v)
		return
	}
	fmt.Printf("  %-12s valor %d: encontrado (no com valor %d)\n", metodo, v, encontrado.Valor)
}

func main() {
	arvore := NovaArvore()
	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	for _, v := range valores {
		arvore.Inserir(v)
	}
	fmt.Printf("Arvore construida com: %v\n", valores)

	fmt.Println("\nBusca por 65:")
	relatar("Buscar", 65, arvore.Buscar(65))
	relatar("BuscarIter", 65, arvore.BuscarIter(65))

	fmt.Println("\nBusca por 45:")
	relatar("Buscar", 45, arvore.Buscar(45))
	relatar("BuscarIter", 45, arvore.BuscarIter(45))

	fmt.Println("\nBusca em arvore vazia:")
	vazia := NovaArvore()
	relatar("Buscar", 10, vazia.Buscar(10))
	relatar("BuscarIter", 10, vazia.BuscarIter(10))
}
