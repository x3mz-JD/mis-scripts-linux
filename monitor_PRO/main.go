package main

import "fmt"

func main() {
    var ping int

    fmt.Println("--- MONITOR CONTINUO (Escribe -1 para salir) ---")

    // Inicio del Bucle Infinito
    for {
        fmt.Print("\nIntroduce el ping actual: ")
        fmt.Scanln(&ping)

        // 1. CONDICIÓN DE SALIDA (Para que no sea eterno)
        if ping == -1 {
            fmt.Println("Cerrando el monitor. ¡Adiós!")
            break // <--- Esto ROMPE el bucle y termina el programa
        }

        // 2. ARREGLANDO EL BUG DEL NEGATIVO (Validación de datos)
        if ping < 0 {
            fmt.Println("❌ Error: El ping no puede ser negativo. Inténtalo de nuevo.")
            continue // <--- Esto salta al principio del bucle (ignora lo de abajo)
        }

        // 3. TU LÓGICA DE AYER
        if ping > 200 {
            fmt.Println("❌ ALERTA: Lag alto.")
        } else if ping > 50 {
            fmt.Println("⚠️ AVISO: Conexión aceptable.")
        } else {
            fmt.Println("✅ ÉXITO: Conexión excelente.")
        }
        // Aquí acaba la vuelta y vuelve a empezar arriba
    }
}
