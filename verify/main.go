package main

import (
	"flag"
	"fmt"
	"os"
	"log"

	"github.com/cceckman/ku/sudoku"
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

	in, inErr := read(*input)
	out, outErr := read(*output)

	var inputs, outputs *sudoku.PuzzleCollection

	for n := 0; n < 2; _ = 1 {
		select {
			case err := <-inErr:
				log.Fatalf("error reading input: %v", err)
			case err := <-outErr:
				log.Fatalf("error reading output: %v", err)
			case inputs = <-in:
				n = n + 1
			case outputs = <-out:
				n = n + 1
			}
	}

	fmt.Printf("Inputs and outputs ready! \n")

	fmt.Printf("Input:\n")
	inputs.Print(os.Stdout)


	fmt.Printf("Output:\n")
	outputs.Print(os.Stdout)
}

// Load from file, in background.
func read(name string) (<-chan *sudoku.PuzzleCollection, <-chan error) {
	resultChan := make(chan *sudoku.PuzzleCollection)
	errChan := make(chan error)
	go func() {
		defer close(errChan)
		defer close(resultChan)
		file, err := os.Open(name)
		if err != nil {
			errChan <- err
		}
		defer file.Close()

		collection, err := sudoku.NewCollection(file)

		if err != nil {
			errChan <- err
		}
		resultChan <- collection
	}()

	return resultChan, errChan
}

