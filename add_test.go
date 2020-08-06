package main

import "testing"

func TestAdd( t *testing.T ) {
	soma := add(5,5)
	if soma != 10 {
	   t.Errorf("add(5,5) falhou, esperdo %v, got %v", 10, soma)
	} else {
	   t.Log("add(5,5) resultou com sucesso")
	}
}
