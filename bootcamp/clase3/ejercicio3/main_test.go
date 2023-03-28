package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Run("salario a", func(t *testing.T) {
		var esperado float32 = 4500
		resultado := calculaSalario(60, "A")
		assert.Equal(t, esperado, resultado, "error")
	})
	t.Run("salario b", func(t *testing.T) {
		var esperado float32 = 1800
		resultado := calculaSalario(60, "B")
		assert.Equal(t, esperado, resultado, "error")
		// 1800 !=  1800.0001
	})
	t.Run("salario c", func(t *testing.T) {
		var esperado float32 = 1000
		resultado := calculaSalario(60, "C")
		assert.Equal(t, esperado, resultado, "error")
	})
}
