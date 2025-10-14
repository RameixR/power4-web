package power4

func Init_Grille(Grille *[6][7]int) {
    for x := 0; x < 6; x++ {
        for y := 0; y < 7; y++ {
            Grille[x][y] = 0
        }
    }
}
