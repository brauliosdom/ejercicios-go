package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Run("promedio", func(t *testing.T) {
		var esperado float32 = 2
		resultado, err := calculaPromedio(1, 2, -1)
		assert.Equal(t, esperado, resultado, err)
	})
}
