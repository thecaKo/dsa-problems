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

func (a *Arvore) PreOrdem() []int {
	valores := []int{}
	var percorrer func(*No)
	percorrer = func(no *No) {
		if no == nil {
			return
		}
		valores = append(valores, no.Valor)
		percorrer(no.Esquerdo)
		percorrer(no.Direito)
	}
	percorrer(a.Raiz)
	return valores
}

// EmOrdem produz valores em ordem crescente porque a propriedade da BST garante
// que toda chave da subarvore esquerda e menor que a raiz e toda chave da
// subarvore direita e maior. Visitando esquerda, raiz e direita nessa ordem, o
// percurso emite primeiro todo o bloco de menores, depois a raiz e por fim todo
// o bloco de maiores. Como a mesma regra vale recursivamente em cada subarvore,
// a concatenacao dos blocos resulta na sequencia ordenada.
func (a *Arvore) EmOrdem() []int {
	valores := []int{}
	var percorrer func(*No)
	percorrer = func(no *No) {
		if no == nil {
			return
		}
		percorrer(no.Esquerdo)
		valores = append(valores, no.Valor)
		percorrer(no.Direito)
	}
	percorrer(a.Raiz)
	return valores
}

func (a *Arvore) PosOrdem() []int {
	valores := []int{}
	var percorrer func(*No)
	percorrer = func(no *No) {
		if no == nil {
			return
		}
		percorrer(no.Esquerdo)
		percorrer(no.Direito)
		valores = append(valores, no.Valor)
	}
	percorrer(a.Raiz)
	return valores
}

func (a *Arvore) EmLargura() []int {
	valores := []int{}
	if a.Raiz == nil {
		return valores
	}
	fila := []*No{a.Raiz}
	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]
		valores = append(valores, atual.Valor)
		if atual.Esquerdo != nil {
			fila = append(fila, atual.Esquerdo)
		}
		if atual.Direito != nil {
			fila = append(fila, atual.Direito)
		}
	}
	return valores
}

func main() {
	arvore := NovaArvore()
	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	for _, v := range valores {
		arvore.Inserir(v)
	}
	fmt.Printf("Arvore construida com: %v\n\n", valores)

	fmt.Printf("Pre-ordem:  %v\n", arvore.PreOrdem())
	fmt.Printf("Em-ordem:   %v\n", arvore.EmOrdem())
	fmt.Printf("Pos-ordem:  %v\n", arvore.PosOrdem())
	fmt.Printf("Em largura: %v\n", arvore.EmLargura())
}
