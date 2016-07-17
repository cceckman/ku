package verify

import (
	"os"
	"strings"
	"testing"

	"github.com/cceckman/ku/puzzle"
)



// Case names where we expect issues.
// TODO as an enhancement to the tests: make it map to []string,
// prefixes of issues that we expect.
type expectedIssues map[string]bool

func getPath(in string) string {
	return os.Getenv("PWD") + string(os.PathSeparator) + in
}

func TestIsSolved(t *testing.T) {
	// TODO: Figure out WTH is going on here. Why doesn't a path implicitly relative to the working
	// directory work? Unfortunately, I'm on a plane and don't have access to the Bazel documentation,
	// so... tag for followup in the devlog.
	cases := map[string]expectedIssues{
		getPath("testdata/suite-a.good.txt"): make(expectedIssues),
		// Case 1: r0 c0 b0 is 0, not 8
		// Case 2: r2 c4 b1 is 6, not 2
		// Case 6: many zeros (partially solved)
		getPath("testdata/suite-a.bad-1.txt"): expectedIssues{
			"Case 1": true,
			"Case 2": true,
			"Case 6": true,
		},
		getPath("testdata/suite-b.good.txt"): make(expectedIssues),
		getPath("testdata/suite-c.good.txt"): make(expectedIssues),
		// Case 1: r5 c1 b5 is A, not 4
		getPath("testdata/suite-c.bad.txt"): expectedIssues{
			"Case 1": true,
		},
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

	t.Logf("Testing file %v\n", path)
	for _, puzzle := range collection.Puzzles {
		// Use the comma ok idiom to gather "expect it to be solved." if there are no expected issues
		_, expectIssues := issuesFor[puzzle.Name]

		solved, issues := IsSolved(puzzle)
		if solved == expectIssues {
			t.Errorf("puzzle %q had unexpected solution state: got: %v expected: %v", puzzle.Name, solved, !expectIssues)
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
