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
	head                Position
	tail                Position
	uniqueTailMovesRepo UniqueTailMovesRepository
}

// NewTailMovementUsecase creates a new TailMovementUsecase.
func NewTailMovementUsecase(repo UniqueTailMovesRepository) TailMovementUsecase {
	return &tailMovementUsecase{
		head:                Position{},
		tail:                Position{},
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
		u.head.X += dx
		u.head.Y += dy
		u.tail = u.moveTail(u.head, u.tail)
	}

	fmt.Println("head:", u.head)
	fmt.Println("tail:", u.tail)
}

// moveTail calculates the new tail position based on head and tail positions.
func (u *tailMovementUsecase) moveTail(head, tail Position) Position {
	// Calculate the absolute difference between head and tail coordinates
	dx := helpers.Abs(head.X - tail.X)
	dy := helpers.Abs(head.Y - tail.Y)

	// Check if the head is two steps directly up, down, left, or right from the tail
	if (dx == 2 && dy == 0) || (dx == 0 && dy == 2) {
		// Move the tail one step towards the head in the same direction
		newPos := Position{X: tail.X + (head.X-tail.X)/2, Y: tail.Y + (head.Y-tail.Y)/2}
		u.uniqueTailMovesRepo.Add(newPos)
		return newPos
	}

	// Check if the head and tail aren't touching and aren't in the same row or column and aren't diagonal already
	if !((dx == 1 && dy == 0) || (dx == 0 && dy == 1) || (head.X == tail.X && head.Y == tail.Y) || (dx == 1 && dy == 1)) {
		// Move the tail one step diagonally towards the head
		newPos := Position{X: tail.X + helpers.Sign(head.X-tail.X), Y: tail.Y + helpers.Sign(head.Y-tail.Y)}
		u.uniqueTailMovesRepo.Add(newPos)
		return newPos
	}

	// If none of the above conditions are met, the tail stays in the same position
	return tail
}

// ResetPosition resets the head and tail positions to (0, 0).
func (u *tailMovementUsecase) ResetPosition() {
	u.head, u.tail = Position{}, Position{}
}

// CountUniqueTailMoves returns the count of unique tail moves.
func (u *tailMovementUsecase) CountUniqueTailMoves() int {
	return u.uniqueTailMovesRepo.Count()
}

// GetTailPosition returns the current tail position.
func (u *tailMovementUsecase) GetTailPosition() Position {
	return u.tail
}

// GetHeadPosition returns the current head position.
func (u *tailMovementUsecase) GetHeadPosition() Position {
	return u.head
}
