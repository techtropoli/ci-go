package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invalido: Resultado %d. Esperado: %d.", total, 30)
	}
}

func TestSubtracao(t *testing.T) {
	total := subtracao(15, 15)

	if total != 0 {
		t.Errorf("Resultado da subtração é invalido: Resultado %d. Esperado: %d.", total, 0)
	}
}

func TestMultiplicacao(t *testing.T) {
	total := multiplicacao(15, 15)

	if total != 225 {
		t.Errorf("Resultado da subtração é invalido: Resultado %d. Esperado: %d.", total, 225)
	}
}
