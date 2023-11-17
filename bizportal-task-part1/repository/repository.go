// Package repository provides an implementation of the UniqueTailMovesRepository interface.
package repository

import . "unique-tail-moves-api/entity"

// UniqueTailMovesRepository defines methods for managing unique tail moves.
type UniqueTailMovesRepository interface {
	Add(Position)
	Count() int
}

// uniqueTailMovesRepository is an implementation of UniqueTailMovesRepository.
type uniqueTailMovesRepository struct {
	positions map[Position]bool
}

// NewUniqueTailMovesRepository creates a new UniqueTailMovesRepository.
func NewUniqueTailMovesRepository() UniqueTailMovesRepository {
	return &uniqueTailMovesRepository{
		positions: make(map[Position]bool),
	}
}

// Add adds a position to the repository.
func (r *uniqueTailMovesRepository) Add(pos Position) {
	r.positions[pos] = true
}

// Count returns the count of unique tail moves.
func (r *uniqueTailMovesRepository) Count() int {
	return len(r.positions)
}
