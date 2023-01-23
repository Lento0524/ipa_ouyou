package main

type flag int

const OK flag = 1
const NG flag = 0
const VISITED flag = 2

func visit(x, y int) {
	maze[x][y] = VISITED
	//stack_visit[stack_top] <- (x,y)
	if x == goal_x && y == goal_y {
		for k := 0; k <= stack_top; k++ {
			//paths[sol_num][k] <-stack_visit[k]
		}
	}
}
