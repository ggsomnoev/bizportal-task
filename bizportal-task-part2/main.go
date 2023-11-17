package main

import (
	"log"
	. "unique-tail-moves-api/controller"
	. "unique-tail-moves-api/repository"
	. "unique-tail-moves-api/usecase"
)

var (
	controller          = NewTailMovementController(usecase)
	usecase             = NewTailMovementUsecase(uniqueTailMovesRepo, 10)
	uniqueTailMovesRepo = NewUniqueTailMovesRepository()
)

func main() {
	fp := "head_movements.txt"
	fi, err := controller.NewFileInput(fp)
	if err != nil {
		log.Fatalf("Couldn't load the provided file: %s", fp)
	}

	// Parse the input moves and perform head/tail movements.
	controller.ParseMovesFromInput(fi)

	log.Printf("Number of unique tail moves: %v\n", controller.CountUniqueTailMoves())
}

// TODO: Run a linter. Possibly more test coverage.
