package main

import "fmt"

// Escreva uma função que receba um ponterio para um inteiro e verifique se esse inteiro é par
// ou ímpar. A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.

func main() {
	var p int = 10
	sumPointer(&p, 10)
	fmt.Println(p)
}
