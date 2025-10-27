package power4

const (
	Empty   = 0
	Player1 = 1
	Player2 = 2
)

func CanPlay(g *[6][7]int, col int) bool {
	return col >= 0 && col < 7 && g[0][col] == Empty
}

func DropToken(g *[6][7]int, col int, player int) (int, bool) {
	if player != Player1 && player != Player2 {
		return -1, false
	}
	if !CanPlay(g, col) {
		return -1, false
	}
	for r := 5; r >= 0; r-- {
		if g[r][col] == Empty {
			g[r][col] = player
			return r, true
		}
	}
	return -1, false
}

func CheckWin(g *[6][7]int, row, col, player int) bool {
	if player == Empty {
		return false
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	for _, d := range dirs {
		if 1+countDir(g, row, col, d[0], d[1], player)+countDir(g, row, col, -d[0], -d[1], player) >= 4 {
			return true
		}
	}
	return false
}

func countDir(g *[6][7]int, r, c, dr, dc, player int) int {
	n := 0
	for {
		r += dr
		c += dc
		if r < 0 || r >= 6 || c < 0 || c >= 7 {
			break
		}
		if g[r][c] != player {
			break
		}
		n++
	}
	return n
}

func IsDraw(g *[6][7]int) bool {
	for c := 0; c < 7; c++ {
		if g[0][c] == Empty {
			return false
		}
	}
	return true
}

func Grille_Jeton(col int, player int, g *[6][7]int) string {
	row, ok := DropToken(g, col, player)
	if !ok {
		if col < 0 || col >= 7 {
			return "Erreur: colonne invalide"
		}
		return "Erreur: colonne pleine"
	}
	if CheckWin(g, row, col, player) {
		return "Fleur placé: victoire"
	}
	if IsDraw(g) {
		return "Fleur placé: match nul"
	}
	return "Fleur placé avec succès"
}
