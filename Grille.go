package power4

import (
	"fmt"
)

func Init_Grille(Grille *[6][7]int){

	for x := 0 ; x < 6; x++ {
        for y := 0; y < 7; y++ {
            Grille[x][y] = 0
		}
	}
	fmt.Print(*Grille)

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