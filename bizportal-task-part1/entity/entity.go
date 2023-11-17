// Package entity defines basic data structure/s.
package entity

import "io"

// Position represents a point in a 2D space with X and Y coordinates.
type Position struct {
	X, Y int
}

// InputSource is an interface extending io.Reader, representing a source of input.
type InputSource interface {
	io.Reader
}
