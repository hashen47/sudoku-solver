# Sudoku Solver


## Usage

1. First install the package 

    ```
    go get github.com/hashen47/sudoku-solver
    ```

2. Import the package 
    ```go
    import (
        "github.com/hashen47/sudoku-solver"
    )
    ```

3. Then pass the input to sudoku_solver.GetSudokuSolutions function. it returns the solutions

    ```go
    package main

    import (
        "fmt"
        "github.com/hashen47/sudoku-solver"
    )

    func main() {
        dataset := sudoku_solver.Board{
            []int{0,0,4,3,0,0,2,0,9,}, // rows
            []int{0,0,5,0,0,9,0,0,1,},
            []int{0,7,0,0,6,0,0,4,3,},
            []int{0,0,6,0,0,2,0,8,7,},
            []int{1,9,0,0,0,7,4,0,0,},
            []int{0,5,0,0,8,3,0,0,0,},
            []int{6,0,0,0,0,0,1,0,5,},
            []int{0,0,3,5,0,8,6,9,0,},
            []int{0,4,2,9,1,0,3,0,0,},
        }

        solutions := sudoku_solver.GetSudokuSolutions(dataset) // []sudoku_solver.Board

        for _, solution := range solutions {
            fmt.Println(solution)
            /*
            8 6 4 3 7 1 2 5 9
            3 2 5 8 4 9 7 6 1
            9 7 1 2 6 5 8 4 3
            4 3 6 1 9 2 5 8 7
            1 9 8 6 5 7 4 3 2
            2 5 7 4 8 3 9 1 6
            6 8 9 7 3 4 1 2 5
            7 1 3 5 2 8 6 9 4
            5 4 2 9 1 6 3 7 8
            */
        }
    }
    ```

