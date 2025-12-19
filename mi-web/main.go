package main

import (
    "fmt"
    "net/http" // Librería para hacer servidores web
)

func main() {
    // Cuando alguien entre a la raíz "/", ejecuta esta función
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>¡Hola Mundo! 🌍</h1><p>Soy un servidor Go corriendo en Docker.</p>")
    })

    fmt.Println("Servidor escuchando en el puerto 8080...")
    // Arranca el servidor
    http.ListenAndServe(":8080", nil)
}
