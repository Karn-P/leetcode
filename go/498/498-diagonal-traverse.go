package main

func findDiagonalOrder(mat [][]int) []int {
    lrow,lcol := len(mat), len(mat[0])
    maxCount := lrow * lcol

    row, col := 0,0
    // 3x3
    // 0,0 0,1 0,2
    // 1,0 1,1 1,2
    // 2,0 2,1 2,2

    // 4x4
    // 0,0 0,1 0,2 0,3
    // 1,0 1,1 1,2 1,3
    // 2,0 2,1 2,2 2,3
    // 3,0 3,1 3,2 3,3

    dir := 1 // 1 = upright | -1 = downleft
    count := 0
    var result []int
    for count < maxCount {

        result = append(result, mat[row][col])
        if dir > 0 {
            // up right
            // hit the top wall
            if row == 0 && col < lcol - 1 {
                col++
                dir = 0 - dir
            // hit the right wall
            } else if col == lcol-1 {
                row++
                dir = 0 - dir
            } else {
                row--
                col++
            }
        } else {
            // move down left
            // hit left wall
            if col == 0 && row < lrow - 1 {
                row++
                dir = 0 - dir
            // hit bottom wall
            } else if row == lrow-1 {
                col++
                dir = 0 - dir
            } else {
                row++
                col--
            }
        }
        count++
    }
    return result
}
