package main

import "fmt"

func main() {
    // 1. Declarar variables (Las cajas donde guardamos info)
    var serverName string
    var ping int

    // 2. Pedir datos al usuario
    fmt.Println("--- MONITOR DE SISTEMAS v1.0 ---")
    
    fmt.Print("Introduce el nombre del servidor: ")
    fmt.Scanln(&serverName) // Escucha lo que escribes y lo guarda en serverName

    fmt.Print("Introduce el ping actual (ms): ")
    fmt.Scanln(&ping) // Escucha y lo guarda en ping

    // 3. Lógica (Tomar decisiones)
    fmt.Println("\nAnalizando el servidor", serverName, "...")

    if ping > 200 {
        fmt.Println("❌ ALERTA: La conexión es muy lenta (Lag alto).")
    } else if ping > 50 {
        fmt.Println("⚠️ AVISO: La conexión es estable, pero podría mejorar.")
    } else {
        fmt.Println("✅ ÉXITO: El servidor vuela. Conexión excelente.")
    }
}
