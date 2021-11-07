package main

type point struct {
	x int
	y int
}

func (p point) add(d point) point {
	return point{
		p.x + d.x,
		p.y + d.y,
	}
}

func (p point) at(grid [][]int) (int, bool) {
	// 上下界線

	if p.x < 0 || p.x >= len(grid) {
		return 0, false
	}
	// 左右界線
	if p.y < 0 || p.y >= len(grid[p.x]) {
		return 0, false
	}

	return grid[p.x][p.y], true
}

func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}
	// 起點
	Q := []point{start}

	// 上下左右
	dirs := []point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, d := range dirs {
			next := cur.add(d)
			val, ok := next.at(maze)

			if val == 1 || !ok {
				continue
			}

			val, ok = next.at(steps)

			if val != 0 || !ok {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.x][next.y] = curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps
}

// func main() {
// 	// 迷宫
// 	maze := [][]int{
// 		{0, 1, 0, 0, 0},
// 		{0, 0, 0, 1, 0},
// 		{0, 1, 0, 1, 0},
// 		{1, 1, 1, 0, 0},
// 		{0, 1, 0, 0, 1},
// 		{0, 1, 0, 0, 0},
// 	}
// 	// 入口
// 	start := point{0, 0}
// 	// 出口
// 	end := point{len(maze) - 1, len(maze[0]) - 1}

// 	steps := walk(maze, start, end)
// 	for _, row := range steps {
// 		for _, val := range row {
// 			// 3位对齐，打印结果
// 			fmt.Printf("%3d", val)
// 		}
// 		fmt.Println()
// 	}
// }
