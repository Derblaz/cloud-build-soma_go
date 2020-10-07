package main

import (
	"testing"
	"DesafioCI_GO/funcs"
)

func TestSoma(t *testing.T){
	a:= 2
	b:= 3
	c:= a+b
	result:= funcs.Soma(a, b)

	if a + b != result {
		t.Errorf("Valor esperado %v, valor retornado %v", c, result )
	}

}