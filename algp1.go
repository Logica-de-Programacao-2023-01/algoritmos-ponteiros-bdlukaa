package main

import "fmt"

// Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n. A função
// deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.

func sumPointer(p *int, n int) {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i
	}
	*p = sum
}

func main() {
	var p int = 10
	sumPointer(&p, 10)
	fmt.Println(p)
}
