package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "power4"
)

func main() {
    var Grille [6][7]int
    power4.Init_Grille(&Grille)

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Template/static"))))

    
    // Page jeux
    http.HandleFunc("/jeux.html", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "Template/jeux.html")
    })

    // Page règles
    http.HandleFunc("/regles.html", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "Template/regles.html")
    })

    // Handler pour jouer
    http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
        col, _ := strconv.Atoi(r.URL.Query().Get("col"))
        player, _ := strconv.Atoi(r.URL.Query().Get("player"))
        msg := power4.Grille_Jeton(col, player, &Grille)
        fmt.Fprintf(w, "<p>%s</p><a href='/'>Retour</a>", msg)
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        http.ServeFile(w, r, "Template/index.html")
    })

    log.Println("Serveur sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}