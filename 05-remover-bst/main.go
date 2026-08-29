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

func (n *No) Minimo() *No {
	if n == nil {
		return nil
	}
	atual := n
	for atual.Esquerdo != nil {
		atual = atual.Esquerdo
	}
	return atual
}

func (n *No) Remover(v int) *No {
	if n == nil {
		return nil
	}

	if v < n.Valor {
		n.Esquerdo = n.Esquerdo.Remover(v)
		return n
	}
	if v > n.Valor {
		n.Direito = n.Direito.Remover(v)
		return n
	}

	if n.Esquerdo == nil && n.Direito == nil {
		return nil
	}
	if n.Esquerdo == nil {
		return n.Direito
	}
	if n.Direito == nil {
		return n.Esquerdo
	}

	sucessor := n.Direito.Minimo()
	n.Valor = sucessor.Valor
	n.Direito = n.Direito.Remover(sucessor.Valor)
	return n
}

func (a *Arvore) Remover(v int) {
	a.Raiz = a.Raiz.Remover(v)
}

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

func ordenado(valores []int) bool {
	for i := 1; i < len(valores); i++ {
		if valores[i-1] >= valores[i] {
			return false
		}
	}
	return true
}

func arvoreExemplo() *Arvore {
	a := NovaArvore()
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		a.Inserir(v)
	}
	return a
}

func testarRemocao(caso string, valor int, detalhe string) {
	a := arvoreExemplo()
	fmt.Printf("%s\n", caso)
	fmt.Printf("  Detalhe:        %s\n", detalhe)
	fmt.Printf("  Antes:          %v\n", a.EmOrdem())

	a.Remover(valor)

	depois := a.EmOrdem()
	fmt.Printf("  Removendo %d\n", valor)
	fmt.Printf("  Depois:         %v\n", depois)
	fmt.Printf("  Raiz:           %d\n", a.Raiz.Valor)
	fmt.Printf("  BST preservada: %t\n\n", ordenado(depois))
}

func main() {
	testarRemocao(
		"Caso 1 - no folha",
		20,
		"o no 20 nao possui filhos, basta desligar o ponteiro do pai",
	)

	testarRemocao(
		"Caso 2 - no com um filho",
		40,
		"o no 40 possui apenas o filho esquerdo 35, que sobe em seu lugar",
	)

	testarRemocao(
		"Caso 3 - no com dois filhos",
		50,
		"o no 50 possui dois filhos; o sucessor em-ordem e 60, menor valor da subarvore direita",
	)

	fmt.Println("Remocao de valor inexistente e de arvore vazia")
	a := arvoreExemplo()
	a.Remover(45)
	fmt.Printf("  Removendo 45 (inexistente): %v\n", a.EmOrdem())
	vazia := NovaArvore()
	vazia.Remover(10)
	fmt.Printf("  Removendo 10 de arvore vazia: %v\n", vazia.EmOrdem())
}
