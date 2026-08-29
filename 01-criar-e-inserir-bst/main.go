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

func (a *Arvore) Contar() int {
	var contar func(*No) int
	contar = func(no *No) int {
		if no == nil {
			return 0
		}
		return 1 + contar(no.Esquerdo) + contar(no.Direito)
	}
	return contar(a.Raiz)
}

func main() {
	arvore := NovaArvore()

	fmt.Println("Arvore recem-criada:")
	fmt.Printf("  Raiz nil: %t\n", arvore.Raiz == nil)
	fmt.Printf("  Em-ordem: %v\n", arvore.EmOrdem())
	fmt.Printf("  Total de nos: %d\n", arvore.Contar())

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	fmt.Printf("\nInserindo os valores: %v\n", valores)
	for _, v := range valores {
		arvore.Inserir(v)
	}
	fmt.Printf("  Em-ordem: %v\n", arvore.EmOrdem())
	fmt.Printf("  Total de nos: %d\n", arvore.Contar())

	duplicados := []int{50, 35, 80}
	fmt.Printf("\nTentando inserir duplicados: %v\n", duplicados)
	for _, v := range duplicados {
		arvore.Inserir(v)
	}
	fmt.Printf("  Em-ordem: %v\n", arvore.EmOrdem())
	fmt.Printf("  Total de nos apos duplicados: %d\n", arvore.Contar())

	vazia := NovaArvore()
	fmt.Println("\nInsercao em uma arvore inicialmente vazia:")
	vazia.Inserir(42)
	fmt.Printf("  Raiz: %d\n", vazia.Raiz.Valor)
	fmt.Printf("  Em-ordem: %v\n", vazia.EmOrdem())
	fmt.Printf("  Total de nos: %d\n", vazia.Contar())
}
