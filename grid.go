package tuiles

import "fmt"

func GenerateGrid(x, y int, edge, void byte) [][]byte {
	grid := make([][]byte, y)
	for height := 0; height < y; height++ {
		grid[height] = make([]byte, x)
		for width := 0; width < x; width++ {
			if height == 0 || height == y-1 || width == 0 || width == x-1 {
				grid[height][width] = edge
			} else {
				grid[height][width] = void
			}
		}
	}
	return grid
}

func PrintGrid(grid [][]byte) {
	for _, g := range grid {
		fmt.Println(string(g) + "\r")
	}
}

func GridByteSelect(grid [][]byte, x, y int, char byte) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == char {
				grid[i][j] = ' '
				break
			}
		}
	}
	grid[x][y] = char
}
