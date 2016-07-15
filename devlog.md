# ku devlog

## 2016-07-12 Tue 20:54
So, this morning, I thought "a Sudoku solver would be a nice little project to
practice Go." In particular, it's a good exercise in concurrency- the
performance will, I expect, be highly tied to how parallel or not the program
can be (by design) or is (by core limitations).

So... also a good exercise in test-driven development! And maybe a good exercise
in note-keeping; so, this devlog to be updated ~live as I walk through this.

(Tonight, I'm watching X-Files while doing this as well... so this won't
necessarily go quickly.)

## 2016-07-12 Tue 20:59
So, first thing's first: let's find some test cases. First link on Google for
"sudoku database" is [here](http://www.menneske.no/sudoku/eng/), which has the
nice property "There is only 1 solution for each puzzle"... except I do want to
support guessing, as that's where the parallelism gets really interesting.

It looks like [this page](http://staffhome.ecm.uwa.edu.au/~00013890/sudokupat.php)
is just about what I'd like... ah, but it's broken.

Oh, hey, but it does make a good observation: "...guaranteed to have a unique
completion, but may (or may not) need backtrackign to solve." That is, "unique
solution" may still mean you need to guess-and-backtrack. (Or,
guess in parallel and discard.)

[Here's](http://english.log-it-ex.com/5.html) a neat article on the space of
Sudoku puzzles, and [here's](http://norvig.com/sudoku.html) one on code /
algorithms- which points to the
[Project Euler](https://projecteuler.net/index.php?section=problems&id=96)
problem with a database! That's a good start; I'll want to put it in
[ICPC](https://icpc.baylor.edu/) form for familiarity (hat tip
@aaronbloomfield.)


## 2016-07-12 Tue 21:47
Yeah, the norgig.com link about is the strategy I was thinking of- constraint
propagation in parallel. Simple enough.

I wrote up an ICPC-like definition in the README. Now to write up a verifier-
"is this a valid solution for the problem"?

## 2016-07-12 Tue 22:00
I'll treat the validator as its own binary- `verify`- with inputs in files
rather than ICPC's stdin / stdout convention. Easy enough to framework that up.

## 2016-07-13 Wed 19:37
Back at the wheel. Grr; OS X has a `screen` without vertical splits. Need to go
get the brew version. Whatever.

Okay, I have loading; time to make some testdata and tests for the verifier.
(Testing of testing of testing...) But that also tests load / print, which is
good.

Hrm, Euler doesn't provide solutions. (I should do more of those, though.) That
[first link](http://menneske.no/sudoku/eng/) does give solutions, and problems,
in HTML; good enough for a few cases, since we want to have a few that are
invalid as well.

## 2016-07-13 Wed 19:46
I realized on the way home that "dimension" is the wrong term for these puzzles-
it's really "size", and the "dimension" is 2. A traditional Su Doku is size 3,
dimension 2; so, numbers 1 to 3^2, with 3^2 boxes, 3^2 in each row / column,
etc. By contrast, a Su Doku of size 2 and dimension 3 would have 2^3 = 8 rows,
columns, and slices (the third dimension), and be subdivided into 2-by-2-by-2
cubes. (It would have 4-by-4-by-4 of these, as a consequence.)

Okay; yeah, the Menneske site works. Copy and paste through `s/\(.\) /\1/g` and
it works fine. Well, okay, a little more:
```
pbpaste | tr ' ' '0' | sed 's/[[:space:]]//g'
```

## 2016-07-14 Thu 10:10
Alright, now actually implementing a bit of the logic- like, indexing, as a
prerequisite for being able to verify. Not doing any optimizations at this
point; I want to make all optimizations tweakable / optional, s.t. we can
basically turn it all off or on at any point.

Of course, there's a drawback there- abstraction / virtualization -> runtime
overhead, which has a performance impact. There's probably a middle ground...
hm, maybe at the build level? Rather than using interfaces (virtualization),
write different concerete types that implement the same interface?

Wait, I said I wasn't doing optimzation yet. :-P So, sure, let's get a verifier
and solver first. (But, yeah, a few different build targets sounds like a not
horrible idea for *this*, even if it's not great generally.)

## 2016-07-14 Thu 11:54
Ran into one of those nice Go features: nondeterministic map iteration.
(Fixed the test with a9282e0.)

I actually really like this feature! It pointed out that I had a
nondeterministic test... well, it did after I ran the test a couple times.

On its applicability here in particular: in README.md, I've said that the
outputs need to have the same labels, but don't have to be in the same order.
Loosening this may be a bit of a premature optimization, in and of itself. The
situation I see that helping is if both of these conditions are true:

1. Some puzzles take much longer than others, and
2. The process is bounded on IO time

If (1) isn't true, then order doesn't matter- we wind up serializing anyway, and
that's the bounding factor. If (2) isn't true, then the fact that we're
serializing doesn't matter.

I think that we can make (2) not true by using `testing.Benchmark` as the core
of the performance measurement. Still writing ICPC-style `solve` and `verify`
binaries to start out with, but ultimately keeping more things in memory and
unmarshalled- not timing the IO- is going to be a better indication of the
things I'm interested in: - the impact of parallelism, function calls,
virtual functions, etc.

## 2016-07-14 Thu 16:26
I note in a comment in 52a39245 - something's weird about the Bazel test
infrastructure, but I'm on a plane without the resources to figure out what. I
get verify_test running, and I can confirm that $PWD/testdata/suite-a.good.txt
exists... but Open("testdata/suite-a.good.txt") doesn't work. I have to
explicitly add in the evaluated "PWD".

Something something $RUNFILES... but there isn't a GO_RUNFILES or RUNFILES in
the environment. Bleh.

## 2016-07-15 Fri 14:49
Yay! Actually have it working- or so it seems.

It's a bit surprising to me that this is as fast as it is- still a fraction of a
second, even though in this first iteration we're completely single-threaded,
brute-force, and not even using the most efficient data structures. Clearly, I
need to scale up the number of puzzles.

This suggests that parallelism may not be a limiting factor- it doesn't take
very long to explore a solution space. Maybe time to start making a list of
variables / optimizations to test against:

- Tail recursion / loopifying- how much overhead does a function call add? (This
  is Go, so probably not much.)
- Parallelism: if we can explore multiple branches at the same time, can we
  speed things up?
  - Test the Go scheduler: do we get win / loss by 
  - Note that this is DFS; parallelism may actually slow things down, because
    ~ all of the non-solution threads will be stealing CPU time from the one
    that will eventually succeed. This is kind of the case with non-parallel
    search too- you have to explore & discard a bunch of duds before getting to
    the right one- so not sure what'll be better.
  - Are there any limits to parallelism that would actuallly help? e.g. barriers
    at a recursion layer? We're only CPU-bounded, not IO bounded, in this
    evaluation, so every thread *can* run to completion without waiting on
    anything else. Another potential limit would be how deeply we do branching-
    e.g. when you just have X empty cells / are Y levels deep / have Z other
    threads running, do your own evaluation serially.
- Heuristic guessing: Evaluate the cells with fewest possibilities first.
  - Are there any cases in which guessing-and-invalidating may be faster than
    running through the logic? That may be a hard question to answer.
  This is probably, actually, a huge improvement that I want to measure.
- Algorithmic improvement to "pure logic" section? (see below)
  - Can we parallelize the evaluation of "what must a cell be"? Yes, I think so;
    is the win worth the cost of thread creation, for that tiny loop?
  - Can we have a more clever serial algorithm, e.g. one based on invalidation
    when something is marked, rather than on marking something, and then looping
    over all the cells *again*? There's something to do with a queue here, e.g.:
    whenever you assign a cell, enqueue all the unknowns in its row / col / box,
    and then the row, col, and box itself. Dequeueing a cell, or dequeueing a
    row / col / box, means doing the 
- Make the problem more difficult
  - Add the ability / option to verify that the solution is unique?

"Pure logic", above- there's two stages to solving a Sudoku.
1. You do some logic, repeatedly:
  1. figuring out whether the other cells in a zone (row / col / box) constrain a
     cell down to one number, 
  2. figure out whether there's only one cell that can be filled by a certain number.
2. If that has stopped making progress, start a new puzzle where one now-empty
   cell is filled in (in a valid way); if that puzzle has a solution (possibly
    after more guesses), then you have a solution with what you filled in.

(2) is most easily parallelized, becase you're evaluating separate puzzles along
different trees of guesses. But we can probably parallelize (1) as well, to some
degree... which may or may not be beneficial.

[This site](http://krazydata.com/hexsudoku) looks like it'll be useful for
sourcing a size-4 Sudoku- that is, 16x16 (hence "hex"). There's some degree of
non-linear scaling with size... not sure what the appropriate function is.
Maybe n^4, for row, col, box, and value? Let's see...

The number of rows, cols, boxes, and values are all N^2 for a size-N (square)
Sudoku; the number of cells is (row * col). The branching factor for guessing is
cells * values- you have to fill each empty cell with each of its possible
values. The logic section generally invovles comparisons of cell * row * col *
box... but that's a little deceptive, because some caching of possibilities,
loop optimzation, and early exiting could make it less.

In any case: I expect that adding more test cases will magnify the differences
in performance between various strategies, but that these puzzles may be too
small to see any substantial difference anyway. But on larger puzzles, it'll
come out more. Let's see- we can go up to size 5, right? Or if we change the
input format, arbitrarily large...

## 2016-07-15 Fri 15:32
So, the next step is to gather a bunch of test cases, and create a benchmark
test.
