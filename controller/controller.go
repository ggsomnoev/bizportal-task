// Package controller handles user input and coordinates the interactions between the usecase and
// other components of this fancy pancy app. The main type is TailMovementController,
// which parses moves from different input sources (files and string inputs for now).
package controller

import (
	"bufio"
	"fmt"
	"os"
	. "unique-tail-moves-api/entity"
	. "unique-tail-moves-api/usecase"
)

// TailMovementController handles user input and controls the tail movement logic.
type TailMovementController struct {
	usecase TailMovementUsecase
}

// NewTailMovementController creates a new TailMovementController.
func NewTailMovementController(usecase TailMovementUsecase) *TailMovementController {
	return &TailMovementController{
		usecase: usecase,
	}
}

// ParseMovesFromInput reads and processes moves from the provided input source.
func (c *TailMovementController) ParseMovesFromInput(inputSource InputSource) error {
	var scanner *bufio.Scanner

	// Reset head and tail positions before parsing new moves.
	c.usecase.ResetPosition()

	// Check if the input source is a file. Helps me with the unit tests.
	if fileInput, ok := inputSource.(*FileInput); ok {
		file, err := os.Open(fileInput.Name())
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	} else {
		// If it's not a file input, create a scanner directly from the input source.
		scanner = bufio.NewScanner(inputSource)
	}

	for scanner.Scan() {
		var movement string
		var steps int
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &movement, &steps)
		if err != nil {
			return fmt.Errorf("error reading input: %v", err)
		}

		// Move the head based on the parsed movement and steps.
		c.usecase.MoveHead(movement, steps)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	return nil
}

// CountUniqueTailMoves returns the count of unique tail moves.
func (c *TailMovementController) CountUniqueTailMoves() int {
	return c.usecase.CountUniqueTailMoves()
}
