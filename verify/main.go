package main

import (
	"flag"
	"fmt"
	"os"
	"log"

	"github.com/cceckman/ku/puzzle"
)

var (
	input  = flag.String("input", "", "Original input file, containing unsolved Sudoku puzzles.")
	output = flag.String("output", "", "Output file from a solver, putatively with solutions to the puzzles in --input.")
	help   = flag.Bool("help", false, "Print a usage message.")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Verifies that the file in output contains solutions to input. It returns nonzero if the solutions are incomplete or invalid.\n")
	}
}


func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(1)
	}

	_, inErr := read(*input)
	_, outErr := read(*output)

	if inErr != nil {
		log.Printf("Error reading input: %v", inErr)
	}
	if outErr != nil {
		log.Printf("Error reading output: %v", outErr)
	}
	if inErr != nil || outErr != nil {
		log.Fatalf("Fatal error!")
	}

	fmt.Printf("Inputs and outputs ready! \n")

}

// Load from file. Don't bother doing it in the background yet.
func read(name string) (*puzzle.Collection, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	collection, err := puzzle.NewCollection(file)

	return collection, err
}

