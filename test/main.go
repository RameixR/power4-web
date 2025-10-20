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
	mu      sync.Mutex
	grid    [6][7]int
	lastMsg string
	tpl     = template.Must(template.ParseFiles(
		"Template/Index.html",
		"Template/jeux.html",
		"Template/regles.html",
	))
)

type PageData struct {
	Grid    [6][7]int
	Message string
	Cols    []int
}

func main() {
	power4.Init_Grille(&grid)

	// Statique (CSS, images…)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Template/static"))))

	// Pages rendues via html/template
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
		Grid:    grid,
		Message: lastMsg,
		Cols:    []int{0, 1, 2, 3, 4, 5, 6},
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
	lastMsg = power4.Grille_Jeton(col, power4.Player1, &grid)
	mu.Unlock()

	http.Redirect(w, r, "/jeux.html", http.StatusSeeOther)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	mu.Lock()
	power4.Init_Grille(&grid)
	lastMsg = "Grille réinitialisée"
	mu.Unlock()
	http.Redirect(w, r, "/jeux.html", http.StatusSeeOther)
}

func render(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tpl.ExecuteTemplate(w, name, data)
}
