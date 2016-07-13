# ku
A Sudoku solver. Don't run this as root.

## Problem statement
From Project Euler:

> Su Doku (Japanese meaning number place) is the name given to a popular puzzle
> concept. Its origin is unclear, but credit must be attributed to
> Leonhard Euler who invented a similar, and much more difficult, puzzle idea
> called Latin Squares. The objective of Su Doku puzzles, however, is to replace
> the blanks (or zeros) in a 9 by 9 grid in such that each row, column, and
> 3 by 3 box contains each of the digits 1 to 9. Below is an example of a
> typical starting puzzle grid and its solution grid.

<table cellpadding="5" cellspacing="0" border="1"><tbody><tr><td
style="font-family:'courier new';font-size:14pt;">0 0 3<br>9 0 0<br>0 0 1</td>
<td style="font-family:'courier new';font-size:14pt;">0 2 0<br>3 0 5<br>8 0
6</td>
<td style="font-family:'courier new';font-size:14pt;">6 0 0<br>0 0 1<br>4 0
0</td>
</tr><tr><td style="font-family:'courier new';font-size:14pt;">0 0 8<br>7 0
0<br>0 0 6</td>
<td style="font-family:'courier new';font-size:14pt;">1 0 2<br>0 0 0<br>7 0
8</td>
<td style="font-family:'courier new';font-size:14pt;">9 0 0<br>0 0 8<br>2 0
0</td>
</tr><tr><td style="font-family:'courier new';font-size:14pt;">0 0 2<br>8 0
0<br>0 0 5</td>
<td style="font-family:'courier new';font-size:14pt;">6 0 9<br>2 0 3<br>0 1
0</td>
<td style="font-family:'courier new';font-size:14pt;">5 0 0<br>0 0 9<br>3 0
0</td>
</tr></tbody></table>

<table cellpadding="5" cellspacing="0" border="1"><tbody><tr><td
style="font-family:'courier new';font-size:14pt;">4 8 3<br>9 6 7<br>2 5 1</td>
<td style="font-family:'courier new';font-size:14pt;">9 2 1<br>3 4 5<br>8 7
6</td>
<td style="font-family:'courier new';font-size:14pt;">6 5 7<br>8 2 1<br>4 9
3</td>
</tr><tr><td style="font-family:'courier new';font-size:14pt;">5 4 8<br>7 2
9<br>1 3 6</td>
<td style="font-family:'courier new';font-size:14pt;">1 3 2<br>5 6 4<br>7 9
8</td>
<td style="font-family:'courier new';font-size:14pt;">9 7 6<br>1 3 8<br>2 4
5</td>
</tr><tr><td style="font-family:'courier new';font-size:14pt;">3 7 2<br>8 1
4<br>6 9 5</td>
<td style="font-family:'courier new';font-size:14pt;">6 8 9<br>2 5 3<br>4 1
7</td>
<td style="font-family:'courier new';font-size:14pt;">5 1 4<br>7 6 9<br>3 8
2</td>
</tr></tbody></table>


> A well constructed Su Doku puzzle has a unique solution and can be solved by
> logic, although it may be necessary to employ "guess and test" methods in
> order to eliminate options (there is much contested opinion over this).
> The complexity of the search determines the difficulty of the puzzle;
> the example above is considered easy because it can be solved by straight
> forward direct deduction.

Su Doku can be generalized to any N<sup>2</sup>-by-N<sup>2</sup> grid,
subdivided into N by N boxes. In a correct solution, each row, column, and box
must contain the same N<sup>2</sup> characters, with no duplicates within a row,
column, or box.

In this problem you will write a Su Doku solver.

### Input
The first line of the input consists of two numbers, separated by a space.
The first number is the dimension of all test cases (N, above).
The second number is the number of test cases; test cases follow on subsequent
lines.

Each test case consists of a header line (giving the test case a name), then N
lines with N characters on each. 1-9 and A-Z are used to indicate filled cells;
0 is used to indicate an empty cell.

### Output
The output should consist of test results. The first line of each case matches
the name of the test case from the input; the next N lines follow the same
format as the input, solved according to the rules of Su Doku.

Test results need not be in the same order as they appeared in the input, but
they must have the same names.
