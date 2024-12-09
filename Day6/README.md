https://habr.com/en/articles/543618/


func spiralOrder(matrix [][]int) (out []int){
if len(matrix) == 0 {
return out
}

n, m := len(matrix), len(matrix[0])

    // processed cells
seen := make([][]bool, n)
for row := 0; row < n; row++ {
seen[row] = make([]bool, m)
}

moves := [][]int{
{0, 1},  // move to the left column
{1, 0},  // move down to the lower row
{0, -1}, // move to the right column
{-1, 0}, // move to the upper row
}

row, col := 0, 0
move := 0

    applyMove := func() {
      seen[row][col] = true
        newRow := row + moves[move][0]
        newCol := col + moves[move][1]
        if newRow == -1 || newRow == n || newCol == -1 || newCol == m || seen[newRow][newCol] {
           // change the direction
           move = (move + 1) % len(moves)      
             row = row + moves[move][0]
           col = col + moves[move][1] 
        } else {
       row, col = newRow, newCol
    }
}

    for i := 0; i < n * m; i++ {
        value := matrix[row][col]
    out = append(out, value)
    row, col = applyMove()
    }

    return out
}