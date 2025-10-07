package power4

const (
	Empty   = 0
	Player1 = 1
	Player2 = 2
)

func DropToken(g *[6][7]int, col int, player int) (int, bool) {
	if col < 0 || col >= 7 {
		return -1, false
	}
	if player != Player1 && player != Player2 {
		return -1, false
	}
	if g[0][col] != Empty {
		return -1, false
	} //la colonne est pleine
	for r := 5; r >= 0; r-- {
		if g[r][col] == Empty {
			g[r][col] = player
			return r, true
		}
	}
	return -1, false
}

func CanPlay(g *[6][7]int, col int) bool {
	return col > 0 && col < 7 && g[0][col] == Empty
}

func CheckWin(g *[6][7]int, row, col, player int) bool {
	if player == Empty {
		return false
	}
	direction := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	for _, d := range direction {
		if 1+countDir(g, row, col, d[0], d[1], player)+countDir(g, row, col, -d[0], -d[1], player) >= 4 {
			return true
		}
	}
	return false
}

func countDir(g *[6][7]int, row, col, dr, dc, player int) int {
	count := 0
	for {
		row += dr
		col += dc 
		if row < 0 || row >= 6 || col < 0 || col >= 7{
			break 
		}
		if g[row][col] != player {
			break
		}
		count ++
	}
	return count
}

func IsDraw (g *[6][7]int) bool {
	for col := 0; col < 7; col++ {
		if g[0][col] == Empty {
		return false
		}
	}
	return true 
}
