package benchmark

import (
	"os"
	"testing"

	"github.com/cceckman/ku/puzzle"
	"github.com/cceckman/ku/solve"
)

func getPath(in string) string {
	return os.Getenv("PWD") + string(os.PathSeparator) + in
}

func fileBenchmark(b *testing.B, fileName string) {
	// setup: load puzzles from suite A
	file, err := os.Open(fileName)
	if err != nil {
		b.Fatal("could not load test suite: ", err)
	}
	defer file.Close()

	collection, err := puzzle.NewCollection(file)
	if err != nil {
		b.Fatal("could not load puzzle collection: ", err)
	}
	b.Logf("Testing file %v\n", fileName)

	b.ResetTimer()
	// Run tests b.N times
	for i := 0; i < b.N; i++ {
		// Yeah, we wind up testing copy performance too.
		cleanCollection := collection.Copy()
		for _, p := range cleanCollection.Puzzles {
			// Don't verify; correctness tested elsewhere.
			if err := solve.Solve(p); err != nil {
				b.Error("error in solving: ", err)
			}
		} // for each puzzle
	} // for each N runs
	// teardown: deferred!
}

func BenchmarkSuite(b *testing.B) {
	fileBenchmark(b, getPath("testdata/suite-a.in.txt"))
}
