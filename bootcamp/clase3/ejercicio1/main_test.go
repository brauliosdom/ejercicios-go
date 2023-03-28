package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Run("sin impuesto", func(t *testing.T) {
		var esperado float32 = 0
		resultado := calculaImpuesto(10000)
		assert.Equal(t, esperado, resultado, "error")
	})
	t.Run("poco impuesto", func(t *testing.T) {
		var esperado float32 = 17000
		resultado := calculaImpuesto(100000)
		assert.Equal(t, esperado, resultado, "error")
	})
	t.Run("con impuesto", func(t *testing.T) {
		var esperado float32 = 270000
		resultado := calculaImpuesto(1000000)
		assert.Equal(t, esperado, resultado, "error")
	})
}
