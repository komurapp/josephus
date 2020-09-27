// Package main apresenta a solução para o Problema de Josephus, usando
// lista ligada circular.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// showHeaders mostra prints para facilitar o debugging.
var showHeaders bool

// node é o elemento da lista, com um valor inteiro e um ponteiro para
// o próximo elemento. Cada nó representa uma pessoa.
type node struct {
	value int
	next  *node
}

// circular é uma lista circular.
// tail é o nó de referência (Head) e len o tamanho da lista.
type circular struct {
	tail *node
	len  int
}

// New retorna uma lista circular com valores-padrão:
// tail=nil e len=0.
func new() *circular {
	return &circular{}
}

// insertHead adiciona um nó na Head da lista, ou seja, no nó
// subsequente à tail.
// Aceita somente inteiros positivos.
func (c *circular) insertHead(element int) {
	if showHeaders {
		fmt.Printf(">>>>>\tInserindo %v\t<<<<<\n", element)
	}
	newHead := &node{
		value: element,
		next:  nil,
	}
	if c.len == 0 {
		c.tail = newHead
		c.tail.next = c.tail
	} else {
		newHead.next = c.tail.next
		c.tail.next = newHead
	}
	c.len++
}

// remove remove o elemento com valor e.
// Equivale a matar a próxima pessoa.
func (c *circular) remove(e int) {
	if showHeaders {
		fmt.Printf(">>>>>\tRemovendo elemento %v\t<<<<<\n", e)
	}
	// Se a lista não tiver nós, não faz nenhuma operação.
	if c.len == 0 {
		return
	}

	pivot := c.tail
	for pivot.next.value != e {
		if pivot.next == c.tail {
			return
		}
		pivot = pivot.next
	}
	pivot.next = pivot.next.next
	c.len--
}

// Display mostra o output da lista criada
func (c *circular) display() {
	if c.len == 0 {
		fmt.Println("[]")
		return
	}

	var list []int
	pivot := c.tail.next
	// Loop até o nó anterior ao tail. Por isso, o valor do pivot
	// deve ser incluído no passo subsequente.
	for i := 1; i < c.len; i++ { // head da lista
		list = append(list, pivot.value)
		pivot = pivot.next
	}
	list = append(list, pivot.value)

	// Para mostrar a lista completa:
	fmt.Println(list)
}

// setupParams pega os argumento para solveJosephus, via stdin.
// Retorna uma lista de lista de int.
func setupParams() [][]int {
	// Número de casos de testes.
	testsuites := 0
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		// Remove os espaços em branco.
		parsedStr := strings.TrimSpace(input.Text())
		// Atoi é equivalente à ParseInt(s, 10, 0).
		if i, err := strconv.Atoi(parsedStr); err == nil {
			testsuites = i
			break
		} else {
			log.Fatalln("Primeiro argumento não é número")
		}
	}

	// Array com tamanho de testsuites, com cada elemento [n, m].
	xtests := make([][]int, testsuites)
	// Chamar "i" vezes o josephus.
	for i := 0; i < testsuites; i++ {
		j := 0
		for input.Scan() {
			xtests[i] = make([]int, 2)
			// Pegar n
			xtests[i][j], _ = strconv.Atoi(input.Text())
			j = 1
			break
		}
		for input.Scan() {
			// Pegar m
			xtests[i][j], _ = strconv.Atoi(input.Text())
			j = 0
			break
		}
	}
	return xtests
}

// solveJosephus cria a lista circular e aplica o algoritmo de remoção.
// n é o tamanho do círculo e m o passo a cada iteração.
// Retorna o sobrevivente.
func solveJosephus(n, m int) int {
	result := 0

	circle := new()
	// [1 2 3 4 5]
	for i := n; i > 0; i-- {
		circle.insertHead(i)
	}
	// [2 3 4 5 1], pois o tail apontará para 1 e o algoritmo seguirá a
	// ordem incremental correta (1 -> 2 -> 3 -> ...).
	circle.tail = circle.tail.next

	for circle.len > 1 {
		prevTail := circle.tail
		for i := 1; i <= m; i++ {
			if i != m {
				circle.tail = circle.tail.next
			}
		}
		// Quando m==c.len, o elemento numa posição p aponta para ele
		// próprio, o que é proibido pelo problema. Então, ele mata o
		// elemento subsequente.
		if circle.tail.next == prevTail {
			circle.tail = circle.tail.next
			circle.remove(circle.tail.next.value)
		} else {
			circle.remove(circle.tail.next.value)
		}
		circle.tail = circle.tail.next
	}

	result = circle.tail.value
	return result
}

// formatJosephusSolution mostra a solução do problema no formato solicitado.
func formatJosephusSolution(args []int, result int) {
	n := args[0]
	m := args[1]

	fmt.Printf("Usando n=%d, m=%d, resultado=%d\n", n, m, result)
}

// main executa o algoritmo para o Problema de Josephus.
func main() {
	args := setupParams()
	//fmt.Println(args)

	xResult := []int{}
	for _, arg := range args {
		xResult = append(xResult, solveJosephus(arg[0], arg[1]))
	}
	//fmt.Println(xResult)

	for i, result := range xResult {
		formatJosephusSolution(args[i], result)
	}
}
