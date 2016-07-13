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
