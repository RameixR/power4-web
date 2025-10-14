package power4

import( 
	"fmt"
)

const (
	Empty   = 0
	Player1 = 1
	Player2 = 2
)

func CanPlay(g *[6][7]int, col int) bool {
	return col >= 0 && col < 7 && g[0][col] == Empty
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
		row += dr //rangé 
		col += dc 
		if row < 0 || row >= 6 || col < 0 || col >= 7 || g[row][col] != player {
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

func Grille_Jeton(y int, nbjoueur int, Grille *[6][7]int) string {
	if y < 0 || y >= 7 {
        return "Erreur: colonne invalide"
    }
    
	for x := 5; x >= 0; x--{
		if Grille[x][y] == 0 {
			Grille[x][y] = nbjoueur
			if CheckWin(Grille, x, y, nbjoueur) {
				return fmt.Sprintf("Le joueur %d a gagné !", nbjoueur)
			} else if IsDraw(Grille) {
				return("Match nul !!")
			} else {
				return( "Jeton placé avec succès")
			}
		}
	}
	fmt.Print(Grille)
	return "Erreur: colonne pleine"
}