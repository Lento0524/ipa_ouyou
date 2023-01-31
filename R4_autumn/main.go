package main

import "fmt"

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
	stack_visit = map[int]point{}
	paths = map[int]map[int]point{}
	visit(start_x, start_y)
	if sol_num == 0 { //オ
		fmt.Print("迷路の解は見つからなかった")
	} else {
		u := sol_num
		for i := 0; i <= u; i++ {
			v := len(paths[i])
			for j := 0; j < v; j++ {
				fmt.Printf("x:%d y:%d\n", paths[i][j].x, paths[i][j].y)
			}
		}
	}
}
