package main

var goal_x, goal_y int
var stack_top int
var sol_num int

func main() {
	stack_top = 0
	sol_num = 0
	setMaze2()
	showMaze()
	start_x, start_y := 1, 1
	goal_x, goal_y = 5, 5
	paths = map[int]map[int]point{}
	visit(start_x, start_y)

}
