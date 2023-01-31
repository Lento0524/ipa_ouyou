package main

func row_ok(n, x int) bool {
	row_top := x - (x % 9) //オ：x-mod(x,9)
	for i := 0; i <= 8; i++ {
		if board[row_top+i] == n { //カ
			return false
		}
	}
	return true
}

func column_ok(n, x int) bool {
	column_top := x % 9 //キ：mod(x,9)
	for i := 0; i <= 8; i++ {
		if board[column_top+i*9] == n { //ク
			return false
		}
	}
	return true
}

func flame_ok(n, x int) bool {
	flame_top := x - (9 * (((x - (x % 9)) / 9) % 3)) - (x % 3) //ケ:(9*mod(div(x,9),3)
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if board[flame_top+9*i+j] == n {
				return false
			}
		}
	}
	return true
}
