package power4

import "fmt"

func Afficher_Grille(g *[6][7]int) string {
	s := "<!doctype html><meta charset='utf-8'><style>td{width:40px;height:40px;text-align:center;border:1px solid #000}</style><table>"
	for i := 0; i < 6; i++ {
		s += "<tr>"
		for j := 0; j < 7; j++ {
			v := g[i][j]
			color := map[int]string{0: "#fff", 1: "#f66", 2: "#66f"}[v]
			s += fmt.Sprintf("<td style='background:%s'>%d</td>", color, v)
		}
		s += "</tr>"
	}
	s += "</table><p>Jouer: /play?col=0..6&player=1|2</p>"
	return s
}
