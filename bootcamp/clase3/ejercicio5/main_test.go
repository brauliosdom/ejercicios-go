package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Run("alimento perro", func(t *testing.T) {
		var esperado float64 = 100
		resultado := animalDog(10)
		assert.Equal(t, esperado, resultado, "error")
	})
	t.Run("alimento gato", func(t *testing.T) {
		var esperado float64 = 50
		resultado := animalCat(10)
		assert.Equal(t, esperado, resultado, "error")
	})
}
