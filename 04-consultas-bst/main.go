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

func (a *Arvore) Altura() int {
	return altura(a.Raiz)
}

func altura(no *No) int {
	if no == nil {
		return -1
	}
	esquerda := altura(no.Esquerdo)
	direita := altura(no.Direito)
	if esquerda > direita {
		return esquerda + 1
	}
	return direita + 1
}

func (a *Arvore) Contar() int {
	return contar(a.Raiz)
}

func contar(no *No) int {
	if no == nil {
		return 0
	}
	return 1 + contar(no.Esquerdo) + contar(no.Direito)
}

func (a *Arvore) ContarFolhas() int {
	return contarFolhas(a.Raiz)
}

func contarFolhas(no *No) int {
	if no == nil {
		return 0
	}
	if no.Esquerdo == nil && no.Direito == nil {
		return 1
	}
	return contarFolhas(no.Esquerdo) + contarFolhas(no.Direito)
}

func (a *Arvore) Minimo() (int, bool) {
	if a.Raiz == nil {
		return 0, false
	}
	atual := a.Raiz
	for atual.Esquerdo != nil {
		atual = atual.Esquerdo
	}
	return atual.Valor, true
}

func (a *Arvore) Maximo() (int, bool) {
	if a.Raiz == nil {
		return 0, false
	}
	atual := a.Raiz
	for atual.Direito != nil {
		atual = atual.Direito
	}
	return atual.Valor, true
}

func relatarExtremo(rotulo string, valor int, ok bool) {
	if !ok {
		fmt.Printf("  %-8s indefinido (arvore vazia)\n", rotulo)
		return
	}
	fmt.Printf("  %-8s %d\n", rotulo, valor)
}

func descrever(titulo string, a *Arvore) {
	fmt.Println(titulo)
	fmt.Printf("  %-8s %d\n", "Altura:", a.Altura())
	fmt.Printf("  %-8s %d\n", "Nos:", a.Contar())
	fmt.Printf("  %-8s %d\n", "Folhas:", a.ContarFolhas())
	minimo, temMinimo := a.Minimo()
	relatarExtremo("Minimo:", minimo, temMinimo)
	maximo, temMaximo := a.Maximo()
	relatarExtremo("Maximo:", maximo, temMaximo)
	fmt.Println()
}

func main() {
	descrever("Arvore vazia (convencao Altura(nil) == -1):", NovaArvore())

	unica := NovaArvore()
	unica.Inserir(42)
	descrever("Arvore com um unico no (convencao Altura(folha) == 0):", unica)

	completa := NovaArvore()
	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	for _, v := range valores {
		completa.Inserir(v)
	}
	descrever(fmt.Sprintf("Arvore construida com %v:", valores), completa)
}
