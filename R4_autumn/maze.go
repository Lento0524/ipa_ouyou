package main

import "fmt"

var maze map[int]map[int]flag

func setMaze1() {
	maze = map[int]map[int]flag{}
	maze[0] = map[int]flag{}
	maze[1] = map[int]flag{}
	maze[2] = map[int]flag{}
	maze[3] = map[int]flag{}
	maze[4] = map[int]flag{}
	maze[5] = map[int]flag{}
	maze[6] = map[int]flag{}
	maze[0][0] = NG
	maze[1][0] = NG
	maze[2][0] = NG
	maze[3][0] = NG
	maze[4][0] = NG
	maze[5][0] = NG
	maze[6][0] = NG
	maze[0][1] = NG
	maze[1][1] = OK
	maze[2][1] = OK
	maze[3][1] = OK
	maze[4][1] = NG
	maze[5][1] = OK
	maze[6][1] = NG
	maze[0][2] = NG
	maze[1][2] = OK
	maze[2][2] = NG
	maze[3][2] = OK
	maze[4][2] = OK
	maze[5][2] = OK
	maze[6][2] = NG
	maze[0][3] = NG
	maze[1][3] = OK
	maze[2][3] = OK
	maze[3][3] = NG
	maze[4][3] = NG
	maze[5][3] = OK
	maze[6][3] = NG
	maze[0][4] = NG
	maze[1][4] = OK
	maze[2][4] = NG
	maze[3][4] = OK
	maze[4][4] = OK
	maze[5][4] = OK
	maze[6][4] = NG
	maze[0][5] = NG
	maze[1][5] = OK
	maze[2][5] = OK
	maze[3][5] = NG
	maze[4][5] = NG
	maze[5][5] = OK
	maze[6][5] = NG
	maze[0][6] = NG
	maze[1][6] = NG
	maze[2][6] = NG
	maze[3][6] = NG
	maze[4][6] = NG
	maze[5][6] = NG
	maze[6][6] = NG
}

func setMaze2() {
	maze = map[int]map[int]flag{}
	maze[0] = map[int]flag{}
	maze[1] = map[int]flag{}
	maze[2] = map[int]flag{}
	maze[3] = map[int]flag{}
	maze[4] = map[int]flag{}
	maze[5] = map[int]flag{}
	maze[6] = map[int]flag{}
	maze[0][0] = NG
	maze[1][0] = NG
	maze[2][0] = NG
	maze[3][0] = NG
	maze[4][0] = NG
	maze[5][0] = NG
	maze[6][0] = NG
	maze[0][1] = NG
	maze[1][1] = OK
	maze[2][1] = OK
	maze[3][1] = OK
	maze[4][1] = NG
	maze[5][1] = OK
	maze[6][1] = NG
	maze[0][2] = NG
	maze[1][2] = OK
	maze[2][2] = NG
	maze[3][2] = OK
	maze[4][2] = OK
	maze[5][2] = OK
	maze[6][2] = NG
	maze[0][3] = NG
	maze[1][3] = OK
	maze[2][3] = OK
	maze[3][3] = OK
	maze[4][3] = NG
	maze[5][3] = OK
	maze[6][3] = NG
	maze[0][4] = NG
	maze[1][4] = OK
	maze[2][4] = NG
	maze[3][4] = OK
	maze[4][4] = OK
	maze[5][4] = OK
	maze[6][4] = NG
	maze[0][5] = NG
	maze[1][5] = OK
	maze[2][5] = OK
	maze[3][5] = NG
	maze[4][5] = NG
	maze[5][5] = OK
	maze[6][5] = NG
	maze[0][6] = NG
	maze[1][6] = NG
	maze[2][6] = NG
	maze[3][6] = NG
	maze[4][6] = NG
	maze[5][6] = NG
	maze[6][6] = NG
}

func showMaze() {
	for y := 6; y >= 0; y-- {
		for x := 0; x <= 6; x++ {
			white := "□"
			brack := "■"
			if maze[x][y] == OK {
				fmt.Print(white, " ")
			} else if maze[x][y] == NG && x == 6 {
				fmt.Println(brack, " ")
			} else if maze[x][y] == NG && x != 6 {
				fmt.Print(brack, " ")
			}
		}
	}
}
