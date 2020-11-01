package main

import "fmt"

/*
	给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地0 表示水域。
	网格中的格子水平和垂直方向相连（对角线方向不相连）。
	整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
*/
//方法一
//func islandPerimeter(grid [][]int) int {
//	var land, border int
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid[0]); j++ {
//			if grid[i][j] == 1 {
//				land++
//				if i < len(grid) - 1 && grid[i+1][j] == 1 {
//					border++
//				}
//				if j < len(grid[0]) - 1 && grid[i][j+1] == 1{
//					border++
//				}
//			}
//
//		}
//	}
//	return 4 * land - 2 * border
//}
//方法二
type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				for _, d := range dir4 {
					if x, y := i+d.x, j+d.y; x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
						ans++
					}
				}
			}
		}
	}
	return
}

func main() {
	grid := [][]int{
		{0,1,0,0},
		{1,1,1,0},
		{0,1,0,0},
		{1,1,0,0},
	}
	fmt.Println(islandPerimeter(grid))
}
