package main

import (
	"html/template"
	"log"
	"net/http"
	"power4"
	"strconv"
	"sync"
)

var (
	mu            sync.Mutex
	grid          [6][7]int
	lastMsg       string
	currentPlayer = power4.Player1
	gameOver      bool
	tpl           = template.Must(template.ParseFiles(
		"Template/Index.html",
		"Template/jeux.html",
		"Template/regles.html",
	))
)

type PageData struct {
	Grid     [6][7]int
	Message  string
	Cols     []int
	Current  int
	GameOver bool
}

func main() {
	power4.Init_Grille(&grid)

	// Statique (CSS, images…)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Template/static"))))

	// Rendue pages html/template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { render(w, "Index.html", nil) })
	http.HandleFunc("/regles.html", func(w http.ResponseWriter, r *http.Request) { render(w, "regles.html", nil) })
	http.HandleFunc("/jeux.html", handleGame)

	// Actions (POST)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/reset", handleReset)

	log.Println("Serveur sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	data := PageData{
		Grid:     grid,
		Message:  lastMsg,
		Cols:     []int{0, 1, 2, 3, 4, 5, 6},
		Current:  currentPlayer,
		GameOver: gameOver,
	}
	lastMsg = ""
	mu.Unlock()
	render(w, "jeux.html", data)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	col, _ := strconv.Atoi(r.FormValue("col"))

	mu.Lock()
	defer mu.Unlock()
	if gameOver {
		lastMsg = "Partie terminée, veuillez cliquer sur réinitialiser"
		http.Redirect(w, r, "/jeux.html", http.StatusSeeOther)
		return
	}

	row, ok := power4.DropToken(&grid, col, currentPlayer)
	if !ok {
		lastMsg = "Colonne pleine"
		http.Redirect(w, r, "/jeux.html", http.StatusSeeOther)
		return
	}

	if power4.CheckWin(&grid, row, col, currentPlayer) {
		if currentPlayer == 1  {
			lastMsg = "Victoire de Marguerite" 
		}else {
			lastMsg = "Victoire de Rose" 
		}
		gameOver = true
		http.Redirect(w, r, "/jeux.html", http.StatusSeeOther)
		return
	}

	if power4.IsDraw(&grid) {
		lastMsg = "Match nul"
		gameOver = true
		http.Redirect(w, r, "/jeux.html", http.StatusSeeOther)
		return
	}

	if currentPlayer == power4.Player1 {
		currentPlayer = power4.Player2
		lastMsg = "A Rose de jouer"
	} else {
		currentPlayer = power4.Player1
		lastMsg = "A Marguerite de jouer"
	}

	http.Redirect(w, r, "/jeux.html", http.StatusSeeOther)

}

func handleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	mu.Lock()
	power4.Init_Grille(&grid)
	currentPlayer = power4.Player1 // FIX: repartir à J1
	gameOver = false               // FIX: débloquer la partie
	lastMsg = "Grille réinitialisée — Joueur 1 commence"
	mu.Unlock()
	http.Redirect(w, r, "/jeux.html", http.StatusSeeOther)
}

func render(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tpl.ExecuteTemplate(w, name, data)
}
