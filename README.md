# Advent of Code

# 2019

I'm going to play with go again this year, definitely not going to win
any speed awards, but I find switching things up from daily
javascript programming to be a nice change.

## Day 01

Wow, hard to believe its been a year.  Good thing I had this repo to
reuse my template.go file and have some structure to start with again.
Easy first puzzle day, as usual, just getting the workflow established
here.

https://github.com/rcy/aoc/tree/master/2019/day01

## Day 02

Most of the time spent working out how to copy arrays

```go
	p := make([]int, len(program))
	copy(p, program)
```

https://github.com/rcy/aoc/tree/master/2019/day02

# 2018

This is the first code I have written using Go!

The aim is not to have the most elegant code, or solve the problems
quickly, but to get working knowledge of the language and to have fun.

## Day 01
https://adventofcode.com/2018/day/1

I wrote the code to part A, and then copied the file over and modified
it for part B.  Going forward I will try to have a single file with
both solutions in it.

## Day 02
https://adventofcode.com/2018/day/2

Worked with arrays.  Didn't use a scanner to process the input file
line by line like I did on day 1, but instead read it all in as bytes
and converted to string, then split into array.

## Day 03
https://adventofcode.com/2018/day/3

Spent a time looking up how maps work in Go.  Lots of for loops in
this language, much different than either Ruby or js where you can
map, filter, and find.

## Day 04
https://adventofcode.com/2018/day/4

Used structs for the first time.  Some things seem tedious, like
finding the max value in a map or hash using loops.  Cheated and
sorted the input using `/bin/sort` (see Makefile in 'inputs/').

## Day 05
https://adventofcode.com/2018/day/5

Messed around with byte arrays.

## Day 06
https://adventofcode.com/2018/day/6

This one was tough, as I kinda got sidetracked building a grid which
wasn't needed.  It did wind up being a neat visualization though.

## Day 07
https://adventofcode.com/2018/day/7

Pointers.  And loops.

## Day 08
https://adventofcode.com/2018/day/8

Early challenges just getting data split and parsed.  Node structure
worked well, off by one error in second part.

## Day 09
https://adventofcode.com/2018/day/9

Solved the first part, then had to rewrite to use linked lists rather
than arrays for second part to fix algorithmic complexity.

## Day 10
https://adventofcode.com/2018/day/10

Briefly considered building a grid, but then decided to print by
scanning.  Part 2 came for free this time.

# 2017

I only did the first 7 days, doing the first three days with javascript, and the last three with elisp.