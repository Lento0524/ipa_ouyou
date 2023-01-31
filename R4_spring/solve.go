package main

import "fmt"

const MAX_BOARD int = 81

func solve(x int) {
	if x > MAX_BOARD-1 {
		print_board()
		fmt.Println("-------------------")
	} else {
		if board[x] != 0 { //ア
			solve(x + 1) //イ
		} else {
			for n := 1; n <= 9; n++ {
				if row_ok(n, x) && column_ok(n, x) && flame_ok(n, x) { //ウ
					board[x] = n
					solve(x + 1)
					board[x] = 0 //エ
				}
			}
		}
	}
}
