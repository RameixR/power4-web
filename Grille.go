package power4

import (
	"fmt"
)

func InitGrille(){
	const ligne, colone = 6, 7
	var grille [ligne][colone]int

	for i := 0 ; i < ligne; i++ {
        for j := 0; j < colone; j++ {
            grille[i][j] = 0
		}
	}
	fmt.Print(grille)
	
}