package main

import (
	"fmt"
	"flag"
	"os"
)

var (
	input = flag.String("input", "", "Original input file, containing unsolved Sudoku puzzles.")
	output = flag.String("output", "", "Output file from a solver, putatively with solutions to the puzzles in --input.")
	help = flag.Bool("help", false, "Print a usage message.")
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

	fmt.Println("vim-go")
}
