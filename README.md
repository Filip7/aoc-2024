# Advent of code 2024 🎄

This advent I'll try to participate in [aoc](https://adventofcode.com/2024).

My first ever.

Wish me luck!

## Initialize a day

To init a day simply run `make init DAY={day number}`,
this will prepare everything needed.

## How to run puzzle solutions[^1]

All commands below are to be run from repo root.

Put your puzzle input into `data.txt` inside `day*/` directories.
Example data is already committed.

### Run with puzzle input

This uses your puzzle input, so if it's missing, there will be all kinds of errors.
Files not included per FAQ of aoc.

Use `make run DAY={number}` and you will get the results. Or run all with `make all`

Example:

```shell
make run DAY=1

...

make run DAY=3

make all
```

### Run with example input

I commit example data since I believe it's the same for everyone.
So no need to add it.

For examples I use go tests (at least for day*/go/ subdirectories)

Use `make test DAY={number}` and you will get the results.
Or run all with `make test-all`

Example:

```shell
make test DAY=1

...

make test DAY=3

make test-all
```

---
[^1]: For me in a month, when I forget everything...
