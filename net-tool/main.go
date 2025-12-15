package main

import (
    "fmt"
    "os"        // Necesario para leer variables de entorno
    "os/exec"
    "strings"
)

func main() {
    // 1. INTENTO LEER LA VARIABLE DESDE FUERA
    target := os.Getenv("TARGET")

    // 2. SEGURIDAD: Si no me pasan nada, uso un valor por defecto
    if target == "" {
        target = "google.com"
        fmt.Println("ℹ️ No se especificó TARGET, usando por defecto:", target)
    }

    fmt.Println("--- PINGEANDO A:", target, "---")

    // El resto es igual que ayer...
    cmd := exec.Command("ping", "-c", "1", target)
    outputBytes, err := cmd.Output()
    resultadoTexto := string(outputBytes)

    if err != nil {
        fmt.Println("❌ Error: El comando falló o no hay internet.")
        return
    }

    if strings.Contains(resultadoTexto, "ttl=") {
        fmt.Println("✅ ÉXITO: El servidor", target, "está vivo.")
    } else {
        fmt.Println("⚠️ OJO: Respuesta extraña.")
    }
} 
