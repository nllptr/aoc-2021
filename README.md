# aoc-2021
Advent of Code 2021

I've done the solutions in a test-driven manner, and the code that processes the actual puzzle input is also inplemented as tests. To see the log messages that gives the answers, you have to run the tests in verbose mode (`go test -v`).

* If you `cd` to a specific day's package directory, you can run the tests for that day with `go test -v`.
* If you want to run everything, cd to the source root and run `go test ./... -v` instead.