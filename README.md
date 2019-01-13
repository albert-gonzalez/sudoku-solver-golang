# Sudoku Solver in Golang

A little example of a sudoku solver in Go. I developed it to learn a little about this language.

## Build & Usage

With docker installed, just run this command inside the project folder:

```sh
docker-compose up
```

Go will generate a binary called `sugoku` into the dist folder. Then, you can run this binary passing the path of the file with the sudoku to solve. For example:

```sh
dist/sugoku example.txt
```

The file must contain a sudoku in this format:

```
0 0 0 0 0 0 6 8 0
0 0 0 0 0 0 0 7 5
0 0 4 1 0 0 0 0 0
0 0 0 0 0 7 0 0 0
0 0 9 0 0 0 1 0 0
0 0 0 0 0 8 0 0 0
8 6 0 9 0 0 0 0 0
0 7 3 0 0 0 0 0 0
0 0 0 0 3 0 5 0 0
```

The 'zero' positions are the empty cells.

## Tests

Run this command to run the tests:

```sh
docker-compose run --rm golang go test
```