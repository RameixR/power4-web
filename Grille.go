package power4

func Init_Grille(Grille *[6][7]int) {
    for x := 0; x < 6; x++ {
        for y := 0; y < 7; y++ {
            Grille[x][y] = 0
        }
    }
}

func Grille_Jeton(y int, nbjoueur int, Grille *[6][7]int) string {
<<<<<<< HEAD
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
=======
    row, ok := DropToken(Grille, y, nbjoueur)
    if !ok {
        if y < 0 || y >= 7 {
            return "Erreur: colonne invalide"
        }
        return "Erreur: colonne pleine"
    }
    if CheckWin(Grille, row, y, nbjoueur) {
        return "Jeton placé: victoire"
    }
    if IsDraw(Grille) {
        return "Jeton placé: match nul"
    }
    return "Jeton placé avec succès"
>>>>>>> 016807378f82ab5723b53ab05aa6fbb8666d84fa
}