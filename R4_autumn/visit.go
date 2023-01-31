package main

type flag int

const OK flag = 1
const NG flag = 0
const VISITED flag = 2

var stack_visit map[int]point
var paths map[int]map[int]point

func visit(x, y int) {
	paths[sol_num] = map[int]point{}
	maze[x][y] = VISITED
	stack_visit[stack_top] = point{
		x: x,
		y: y,
	}
	if x == goal_x && y == goal_y {
		for k := 0; k <= stack_top; k++ {
			paths[sol_num][k] = stack_visit[k]
		}
		sol_num += 1
	} else {
		stack_top += 1
		if maze[x][y+1] == OK {
			visit(x, y+1)
		}
		if maze[x+1][y] == OK {
			visit(x+1, y)
		}
		if maze[x][y-1] == OK {
			visit(x, y-1)
		}
		if maze[x-1][y] == OK {
			visit(x-1, y)
		}
		stack_top -= 1 //ウ
	}
	maze[x][y] = OK //エ

}

type point struct {
	x int
	y int
}
