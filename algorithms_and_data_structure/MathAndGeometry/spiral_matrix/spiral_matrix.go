package main

import "fmt"

func spiralMatrix(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    
    result := []int{}
    top, bottom := 0, len(matrix)-1
    left, right := 0, len(matrix[0])-1

    for top <= bottom && left <= right {
        // Move left to right along top row
        for i := left; i <= right; i++ {
            result = append(result, matrix[top][i])
        }
        top++

        // Move top to bottom along right column
        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][right])
        }
        right--

        // Move right to left along bottom row, if still valid
        if top <= bottom {
            for i := right; i >= left; i-- {
                result = append(result, matrix[bottom][i])
            }
            bottom--
        }

        // Move bottom to top along left column, if still valid
        if left <= right {
            for i := bottom; i >= top; i-- {
                result = append(result, matrix[i][left])
            }
            left++
        }
    }
    return result
}

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println(spiralMatrix(matrix)) // Output: [1 2 3 6 9 8 7 4 5]
}
