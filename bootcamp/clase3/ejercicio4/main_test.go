package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Run("minimo", func(t *testing.T) {
		var esperado float32 = 1
		var resultado float32 = minFunc(1, 2, 3)
		assert.Equal(t, esperado, resultado, "error")
	})
	t.Run("promedio", func(t *testing.T) {
		var esperado float32 = 2
		var resultado float32 = averageFunc(1, 2, 3)
		assert.Equal(t, esperado, resultado, "error")
	})
	t.Run("maximo", func(t *testing.T) {
		var esperado float32 = 3
		var resultado float32 = maxFunc(1, 2, 3)
		assert.Equal(t, esperado, resultado, "error")
	})
}
