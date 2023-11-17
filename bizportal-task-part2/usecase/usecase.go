// Package usecase defines the business logic. (moving the head and tail, counting the unique tail moves)
package usecase

import (
	"fmt"
	"log"
	. "unique-tail-moves-api/entity"
	"unique-tail-moves-api/helpers"
	. "unique-tail-moves-api/repository"
)

// TailMovementUsecase defines methods for tail movement logic.
type TailMovementUsecase interface {
	MoveHead(movement string, steps int)
	ResetPosition()
	GetHeadPosition() Position
	GetTailPosition() Position
	CountUniqueTailMoves() int
}

// tailMovementUsecase is an implementation of TailMovementUsecase.
type tailMovementUsecase struct {
	positions           []Position
	uniqueTailMovesRepo UniqueTailMovesRepository
}

// NewTailMovementUsecase creates a new TailMovementUsecase.
func NewTailMovementUsecase(repo UniqueTailMovesRepository, size int) TailMovementUsecase {
	return &tailMovementUsecase{
		positions:           make([]Position, size),
		uniqueTailMovesRepo: repo,
	}
}

// MoveHead moves the head according to the provided movement and steps.
func (u *tailMovementUsecase) MoveHead(movement string, steps int) {
	var dx, dy int

	switch movement {
	case "L":
		dx = -1
	case "R":
		dx = 1
	case "D":
		dy = -1
	case "U":
		dy = 1
	default:
		log.Fatalf("Invalid movement %v\n", movement)
		return
	}

	fmt.Printf("Head Moving %s %v times\n", movement, steps)
	for i := 0; i < steps; i++ {
		u.positions[0].X += dx
		u.positions[0].Y += dy
		u.moveTail()
	}
	fmt.Println("coordinates:", u.positions)
}

// moveTail moves rest of the body/tail
func (u *tailMovementUsecase) moveTail() {
	for id := 1; id < len(u.positions); id++ {
		prevID := id - 1
		prevTail := u.positions[prevID]
		u.positions[id] = u.moveSegment(u.positions[id], prevTail)
	}

	u.uniqueTailMovesRepo.Add(u.positions[len(u.positions)-1])
}

// moveSegment calculates the new tail position based on head and tail positions.
func (u *tailMovementUsecase) moveSegment(tail, prevTail Position) Position {
	// Calculate the absolute difference between head and tail coordinates
	dx := helpers.Abs(prevTail.X - tail.X)
	dy := helpers.Abs(prevTail.Y - tail.Y)

	// Check if the head is two steps directly up, down, left, or right from the tail
	if (dx == 2 && dy == 0) || (dx == 0 && dy == 2) {
		// Move the tail one step towards the head in the same direction
		return Position{X: tail.X + (prevTail.X-tail.X)/2, Y: tail.Y + (prevTail.Y-tail.Y)/2}
	}

	// Check if the head and tail aren't touching and aren't in the same row or column and aren't diagonal already
	if !((dx == 1 && dy == 0) || (dx == 0 && dy == 1) || (prevTail.X == tail.X && prevTail.Y == tail.Y) || (dx == 1 && dy == 1)) {
		// Move the tail one step diagonally towards the head
		return Position{X: tail.X + helpers.Sign(prevTail.X-tail.X), Y: tail.Y + helpers.Sign(prevTail.Y-tail.Y)}
	}

	// If none of the above conditions are met, the tail stays in the same position
	return tail
}

// ResetPosition resets the head and tail(body) positions to (0, 0).
func (u *tailMovementUsecase) ResetPosition() {
	for i := range u.positions {
		u.positions[i] = Position{X: 0, Y: 0}
	}
}

// CountUniqueTailMoves returns the count of unique tail moves. The last element is considered tail.
func (u *tailMovementUsecase) CountUniqueTailMoves() int {
	return u.uniqueTailMovesRepo.Count()
}

// GetHeadPosition returns the position of the head.
func (u *tailMovementUsecase) GetHeadPosition() Position {
	if len(u.positions) == 0 {
		return Position{}
	}
	return u.positions[0]
}

// GetTailPosition returns the position of the tail.
func (u *tailMovementUsecase) GetTailPosition() Position {
	if len(u.positions) < 2 {
		return Position{}
	}
	return u.positions[len(u.positions)-1]
}
