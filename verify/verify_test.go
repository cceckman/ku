package verify

import (
	"os"
	"strings"
	"testing"

	"github.com/cceckman/ku/puzzle"
)

const (
	goodFile = "testdata/suite-a.good.txt"
	badFile  = "testdata/suite-a.bad-1.txt"
)

// Case names where we expect issues.
// TODO as an enhancement to the tests: make it map to []string,
// prefixes of issues that we expect.
type expectedIssues map[string]bool

func TestIsSolved(t *testing.T) {
	// TODO: Figure out WTH is going on here. Why doesn't a path implicitly relative to the working
	// directory work? Unfortunately, I'm on a plane and don't have access to the Bazel documentation,
	// so... tag for followup in the devlog.
	cases := map[string]expectedIssues{
		os.Getenv("PWD") + string(os.PathSeparator) + goodFile: make(expectedIssues),
		os.Getenv("PWD") + string(os.PathSeparator) + badFile:  make(expectedIssues),
	}

	for k, v := range cases {
		testIsSolved(t, k, v)
	}
}

func testIsSolved(t *testing.T, path string, issuesFor expectedIssues) {

	// Setup: load the file, which may have some bad cases.
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("could not open testdata file: %v", err)
	}
	defer file.Close()
	collection, err := puzzle.NewCollection(file)
	if err != nil {
		t.Fatalf("could not load puzzle collection: %v", err)
	}

	for _, puzzle := range collection.Puzzles {
		// Use the comma ok idiom to gather "expect it to be solved." if there are no expected issues
		_, expectSolved := issuesFor[puzzle.Name]

		solved, issues := IsSolved(puzzle)
		if solved != expectSolved {
			t.Errorf("puzzle %q had unexpected solution state: got: %v expected: %v", puzzle.Name, solved, expectSolved)
			if len(issues) != 0 {
				t.Errorf("got issues:\n%v", strings.Join(issues, "\n"))
			}
			// TODO include expected issues as well
		} else {
			// OK; just log that it met expectations.
			t.Logf("puzzle %q met expectations; detected issues: %v", puzzle.Name, strings.Join(issues, "\n"))
		}
	}
}
