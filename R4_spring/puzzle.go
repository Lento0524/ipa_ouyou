package main

import "fmt"

var board map[int]int

func setPuzzle() {
	board = map[int]int{}
	board[0] = 2
	board[1] = 0
	board[2] = 1
	board[3] = 0
	board[4] = 9
	board[5] = 0
	board[6] = 7
	board[7] = 0
	board[8] = 0

	board[9] = 0
	board[10] = 4
	board[11] = 0
	board[12] = 2
	board[13] = 0
	board[14] = 0
	board[15] = 3
	board[16] = 0
	board[17] = 0

	board[18] = 5
	board[19] = 0
	board[20] = 0
	board[21] = 0
	board[22] = 0
	board[23] = 8
	board[24] = 0
	board[25] = 2
	board[26] = 9

	board[27] = 0
	board[28] = 9
	board[29] = 0
	board[30] = 6
	board[31] = 7
	board[32] = 0
	board[33] = 2
	board[34] = 0
	board[35] = 0

	board[36] = 6
	board[37] = 0
	board[38] = 0
	board[39] = 3
	board[40] = 0
	board[41] = 5
	board[42] = 0
	board[43] = 0
	board[44] = 4

	board[45] = 0
	board[46] = 0
	board[47] = 7
	board[48] = 0
	board[49] = 4
	board[50] = 9
	board[51] = 0
	board[52] = 1
	board[53] = 0

	board[54] = 7
	board[55] = 6
	board[56] = 0
	board[57] = 9
	board[58] = 0
	board[59] = 0
	board[60] = 0
	board[61] = 0
	board[62] = 3

	board[63] = 0
	board[64] = 0
	board[65] = 9
	board[66] = 0
	board[67] = 0
	board[68] = 6
	board[69] = 0
	board[70] = 4
	board[71] = 0

	board[72] = 0
	board[73] = 0
	board[74] = 4
	board[75] = 0
	board[76] = 1
	board[77] = 0
	board[78] = 6
	board[79] = 0
	board[80] = 0

}

func print_board() {
	for i := 0; i <= 80; i++ {
		if i%9 == 8 {
			fmt.Println(board[i], " ")
		} else {
			fmt.Print(board[i], " ")
		}

	}
}
