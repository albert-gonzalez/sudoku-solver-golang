package main
import "fmt"
import "errors"
import "bufio"
import "os"
import "strings"
import "strconv"

const SUDOKU_LENGTH = 9
const PARTIAL_SQUARE_LENGTH = 3

type sudoku struct {
    cells [][]int
}

type position struct {
    row int
    column int
}

func main() {
    input, err := parseInput(os.Args)
    var solvedSudoku sudoku

    if err == nil {
        solvedSudoku, err = solveSudoku(input)
    }

    printResult(solvedSudoku, err)
}

func parseInput(args []string) (sudoku, error) {
    sudo := sudoku{[][]int{}}

    if len(args) != 2 {
        return sudo, errors.New("Invalid input. Usage: sugoku pathfile")
    }

    file, err := openFile(args)
    defer file.Close()

    if err != nil {
        return sudo, err
    }

    return transformFileToSudoku(file)
}

func openFile(args []string) (*os.File, error){
    path := args[1]
    file, err := os.Open(path)
    
    return file, err
}

func transformFileToSudoku(file *os.File) (sudoku, error) {
    sudo := sudoku{[][]int{}}
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        row, err := sliceStringToInt(strings.Split(scanner.Text(), " "))

        if err != nil {
            return sudo, err
        }

        sudo.cells = append(sudo.cells, row)
    }

    if err := scanner.Err(); err != nil {
        return sudo, err
    }

    return sudo, nil
}

func sliceStringToInt(stringArray []string) ([]int, error) {
    stringInt := make([]int, 0, len(stringArray))
    for _, stringValue := range stringArray {
        intValue, err := strconv.Atoi(stringValue)
        if err != nil {
            return stringInt, err
        }
        stringInt = append(stringInt, intValue)
    }
    return stringInt, nil
}

func solveSudoku(sudo sudoku) (sudoku, error) {
    err := checkSudoku(sudo)

    if err != nil {
        return sudo, err
    }

    positions := getFreePositions(sudo)

    sudo, err = fillSudokuCells(sudo, positions)

    return sudo, err
}

func getFreePositions(sudo sudoku) []position {
    positions := []position{}

    for row := 0; row < len(sudo.cells); row++ {
        for column := 0; column < len(sudo.cells[row]); column++ {
            if sudo.cells[row][column] == 0 {
                positions = append(positions, position{row, column})
            }
        }
    }

    return positions
}

func fillSudokuCells(sudo sudoku, positions []position) (sudoku, error) {
    index := 0

    for index = 0; index < len(positions) && index >= 0; index++ {
        pos := positions[index]
        cell := &sudo.cells[pos.row][pos.column]
        isValid := false

        for !isValid {
            *cell++
            isValid = checkSudokuCell(sudo, pos)
        }

        if *cell > SUDOKU_LENGTH {
            *cell = 0
            index = getPreviousIndex(index)
        }
    }

    if index < 0 {
        return sudo, errors.New("Sudoku without solution")    
    }

    return sudo, nil
}

func getPreviousIndex(index int) int {
    return index - 2
}

func checkSudoku(sudo sudoku) error {
    for row := 0; row < len(sudo.cells); row++ {
        for column := 0; column < len(sudo.cells[row]); column++ {
            cell := sudo.cells[row][column]
            if cell != 0 && !checkSudokuCell(sudo, position{row, column}) {
                return errors.New("Invalid Input")
            }
        }
    }

    return nil
}

func checkSudokuCell(sudo sudoku, pos position) bool {
    return checkRow(sudo, pos) && checkColumn(sudo, pos) && checkSquare(sudo, pos)
}

func checkRow(sudo sudoku, pos position) bool {
    cell := sudo.cells[pos.row][pos.column]
    
    count := 0
    for column := 0; column < len(sudo.cells[pos.row]); column++ {
        if (cell == sudo.cells[pos.row][column]) {
            count++
        }
    }

    return count == 1
}

func checkColumn(sudo sudoku, pos position) bool {
    cell := sudo.cells[pos.row][pos.column]
    count := 0
    
    for row := 0; row < len(sudo.cells); row++ {
        if (cell == sudo.cells[row][pos.column]) {
            count++
        }
    }

    return count == 1
}

func checkSquare(sudo sudoku, pos position) bool {
    cell := sudo.cells[pos.row][pos.column]
    square := getCurrentSquare(sudo, pos)

    count := 0
    for row := square.row * PARTIAL_SQUARE_LENGTH; row < (square.row + 1) * PARTIAL_SQUARE_LENGTH; row++ {
        for column := square.column * PARTIAL_SQUARE_LENGTH; column < (square.column + 1) * PARTIAL_SQUARE_LENGTH; column++ {
            if (cell == sudo.cells[row][column]) {
                count++
            }    
        }
    }

    return count == 1
}

func getCurrentSquare(sudo sudoku, pos position) position {
    row := pos.row / PARTIAL_SQUARE_LENGTH
    column := pos.column / PARTIAL_SQUARE_LENGTH

    return position{row, column}
}

func printResult(solvedSudoku sudoku, err error) {
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, row := range solvedSudoku.cells {
        fmt.Println(row)
    }
}