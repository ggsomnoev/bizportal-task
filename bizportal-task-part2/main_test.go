// Package main_test contains test functions for the main package.
package main

import (
	"strings"
	"testing"

	. "unique-tail-moves-api/entity"

	"github.com/stretchr/testify/assert"
)

// TestCountUniqueTailMoves tests the accuracy of the CountUniqueTailMoves functionality
func TestCountUniqueTailMoves(t *testing.T) {
	input := "L 3\nR 2\nU 4\nD 1\n"

	// Parse the input moves and perform tail movements.
	err := controller.ParseMovesFromInput(strings.NewReader(input))
	assert.NoError(t, err)

	assert.Equal(t, 1, controller.CountUniqueTailMoves())
}

// TestHeadTailPositions tests the accuracy of ParseMovesFromInput functionality
func TestHeadTailPositions(t *testing.T) {
	input := "R 1\nU 2\n"

	// Parse the input moves and perform tail movements.
	err := controller.ParseMovesFromInput(strings.NewReader(input))
	assert.NoError(t, err)

	expectedHeadPos := Position{X: 1, Y: 2}
	expectedTailPos := Position{X: 0, Y: 0}

	// Check if the current head and tail positions match the expected values.
	assert.Equal(t, expectedHeadPos, usecase.GetHeadPosition())
	assert.Equal(t, expectedTailPos, usecase.GetTailPosition())
}
