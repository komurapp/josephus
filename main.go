// Package main apresenta a solução para o Problema de Josephus, usando
// lista ligada circular.
package main

import "fmt"

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

	// Mostra cada elemento da lista.
	//for _, e := range list {
	//fmt.Println(e)
	//}
	// Para mostrar a lista completa:
	fmt.Println(list)
}

func main() {
	showHeaders = true
	fmt.Println("solveJosephus(n, m)")
	c := new()
	n := 10

	if showHeaders {
		fmt.Println("<<<<<<<<<< Circular Linked List >>>>>>>>>>")
		for i := 0; i < n; i++ {
			c.insertHead(i)
		}
		c.remove(2)
		c.display()
		len := c.len
		fmt.Printf("len: %d\n", len)
	}

}
