package power4

import (
	"fmt"
)

func Init_Grille(){

	for x := 0 ; x < 7; x++ {
        for y := 0; y < 6; y++ {
            *Grille[x][y] = 0
		}
	}
	fmt.Print(*Grille)

}

func Grille_Jeton(y int, nbjoueur int){
	for x := 5; x > 0; x--{
		if *Grille[y][x] != 0 {
			break
		} else {
			*Grille[x][y] = nbjoueur
		}
	}
	return *Grille
}