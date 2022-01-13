package main

/*A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and space is marked as 1 and 0 respectively in the grid.*/

//walk from every point.  if we reach the end, count it
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    var valid int
    var walk func([][]int, int, int)bool

    walk = func(g [][]int, row, col int)bool{
        fmt.Println(row,col)
        if row < 0 || col < 0 || row >= len(g) || col >= len(g[0]){
            return false
        }
        if g[row][col] == 1{
            return false
        }
        if row == len(g)-1 && col == len(g[0])-1{
            valid++
        }
        return walk(g, row+1, col) || walk(g,row,col+1)
    }

    walk(obstacleGrid, 0,0)
    return valid
}
