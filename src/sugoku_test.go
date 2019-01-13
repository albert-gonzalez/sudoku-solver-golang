package main

import "testing"

func TestItShouldSolveASudokuWhenTheSudokuHasSolution(t *testing.T) {

	inputSudokus := getInputSudokus()
	expectedSudokus := getExpectedSudokus()

	for i, inputSudoku := range inputSudokus {
		solvedSudoku, _ := solveSudoku(inputSudoku)
		expectedSudoku := expectedSudokus[i]

		if (!areEqual(expectedSudoku, solvedSudoku)) {
			t.Error("Sudoku was solved wrongly", "expected", expectedSudoku, "got", solvedSudoku)
		}
	}
}

func TestItShouldReturnANoSolutionErrorWhenTheSudokuDoesNotHaveSolution(t *testing.T) {

	_, err := solveSudoku(sudoku{[][]int{
	  {2, 0, 5, 3, 8, 0, 9, 0, 0},
	  {0, 0, 0, 1, 4, 6, 0, 0, 5},
	  {0, 0, 1, 0, 0, 5, 0, 0, 0},
	  {1, 8, 0, 0, 0, 0, 0, 3, 7},
	  {0, 7, 3, 8, 0, 4, 2, 9, 6},
	  {6, 2, 0, 0, 5, 0, 4, 1, 0},
	  {0, 3, 0, 5, 9, 8, 1, 0, 4},
	  {0, 1, 0, 0, 2, 0, 0, 5, 9},
	  {9, 5, 7, 4, 0, 0, 0, 6, 0},
	}})
  
	if err == nil || err.Error() != "Sudoku without solution" {
		t.Error("Not returned the error Sudoku without solution")
	}
}

func TestItShouldReturnAInvalidInputErrorWhenTheInputIsWrong(t *testing.T) {

	_, err := solveSudoku(sudoku{[][]int{
	  {1, 0, 5, 3, 8, 0, 9, 0, 0},
	  {0, 0, 0, 1, 4, 6, 0, 0, 5},
	  {0, 0, 1, 0, 0, 5, 0, 0, 0},
	  {1, 8, 0, 0, 0, 0, 0, 3, 7},
	  {0, 7, 3, 8, 0, 4, 2, 9, 6},
	  {6, 2, 0, 0, 5, 0, 4, 1, 0},
	  {0, 3, 0, 5, 9, 8, 1, 0, 4},
	  {0, 1, 0, 0, 2, 0, 0, 5, 9},
	  {9, 5, 7, 4, 0, 0, 0, 6, 0},
	}})
  
	if err == nil || err.Error() != "Invalid Input" {
		t.Error("Not returned the error Invalid Input")
	}
}

func areEqual(firstSudoku, secondSudoku sudoku) bool {
	firstCells := firstSudoku.cells
	secondCells := secondSudoku.cells
    if len(firstCells) != len(secondCells) {
        return false
	}
	
    for i, row := range firstCells {
		for j, cell := range row {
			if cell != secondCells[i][j] {
				return false
			}
		}
	}
	
    return true
}

func getInputSudokus() []sudoku {
	return []sudoku{
		{[][]int{
			{0, 0, 5, 3, 8, 0, 9, 0, 0},
			{0, 0, 0, 1, 4, 6, 0, 0, 5},
			{0, 0, 1, 0, 0, 5, 0, 0, 0},
			{1, 8, 0, 0, 0, 0, 0, 3, 7},
			{0, 7, 3, 8, 0, 4, 2, 9, 6},
			{6, 2, 0, 0, 5, 0, 4, 1, 0},
			{0, 3, 0, 5, 9, 8, 1, 0, 4},
			{0, 1, 0, 0, 2, 0, 0, 5, 9},
			{9, 5, 7, 4, 0, 0, 0, 6, 0},
		}},
		{[][]int{
			{0, 1, 0, 0, 0, 6, 0, 0, 0},
			{7, 0, 0, 0, 1, 4, 0, 0, 0},
			{0, 0, 0, 7, 0, 2, 1, 0, 0},
			{8, 0, 0, 0, 0, 3, 7, 0, 5},
			{2, 0, 7, 0, 8, 0, 0, 0, 0},
			{6, 3, 4, 0, 7, 5, 8, 1, 9},
			{0, 7, 9, 1, 0, 0, 0, 0, 0},
			{0, 8, 0, 0, 4, 0, 0, 9, 1},
			{0, 4, 0, 5, 0, 9, 3, 8, 7},
		}},
		{[][]int{
			{0, 0, 0, 0, 0, 0, 6, 8, 0},
			{0, 0, 0, 0, 0, 0, 0, 7, 5},
			{0, 0, 4, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 7, 0, 0, 0},
			{0, 0, 9, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 8, 0, 0, 0},
			{8, 6, 0, 9, 0, 0, 0, 0, 0},
			{0, 7, 3, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 3, 0, 5, 0, 0},
		}},
	}
}

func getExpectedSudokus() []sudoku {
	return []sudoku{
		{[][]int {
			{7, 6, 5, 3, 8, 2, 9, 4, 1},
			{3, 9, 2, 1, 4, 6, 7, 8, 5},
			{8, 4, 1, 9, 7, 5, 6, 2, 3},
			{1, 8, 4, 2, 6, 9, 5, 3, 7},
			{5, 7, 3, 8, 1, 4, 2, 9, 6},
			{6, 2, 9, 7, 5, 3, 4, 1, 8},
			{2, 3, 6, 5, 9, 8, 1, 7, 4},
			{4, 1, 8, 6, 2, 7, 3, 5, 9},
			{9, 5, 7, 4, 3, 1, 8, 6, 2},
		}},
		{[][]int{
			{4, 1, 5, 3, 9, 6, 2, 7, 8},
			{7, 2, 3, 8, 1, 4, 9, 5, 6},
			{9, 6, 8, 7, 5, 2, 1, 3, 4},
			{8, 9, 1, 4, 6, 3, 7, 2, 5},
			{2, 5, 7, 9, 8, 1, 4, 6, 3},
			{6, 3, 4, 2, 7, 5, 8, 1, 9},
			{5, 7, 9, 1, 3, 8, 6, 4, 2},
			{3, 8, 2, 6, 4, 7, 5, 9, 1},
			{1, 4, 6, 5, 2, 9, 3, 8, 7},
		}},
		{[][]int{
			{3, 9, 2, 5, 7, 4, 6, 8, 1},
			{6, 1, 8, 3, 2, 9, 4, 7, 5},
			{7, 5, 4, 1, 8, 6, 2, 9, 3},
			{5, 3, 6, 2, 1, 7, 8, 4, 9},
			{4, 8, 9, 6, 5, 3, 1, 2, 7},
			{1, 2, 7, 4, 9, 8, 3, 5, 6},
			{8, 6, 5, 9, 4, 1, 7, 3, 2},
			{2, 7, 3, 8, 6, 5, 9, 1, 4},
			{9, 4, 1, 7, 3, 2, 5, 6, 8},
		}},
	}
}