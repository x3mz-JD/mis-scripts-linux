package main

import (
    "fmt"
    "os"
    "os/exec" // <--- Paquete para ejecutar comandos de Linux
)

func main() {
    fmt.Println("--- LANZANDO PING A GOOGLE ---")

    // 1. Preparamos el comando (Como si lo escribieras en la terminal)
    // "ping" es el programa, "-c" y "1" son los argumentos (solo 1 paquete)
    cmd := exec.Command("ping", "-c", "1", "google.com")

    // 2. Conectamos la salida del comando a TU pantalla
    // Si no hacemos esto, el ping ocurre "en silencio" y no verías nada.
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // 3. ¡Fuego! Ejecutamos el comando
    err := cmd.Run()

    // 4. Verificamos qué pasó
    if err != nil {
        fmt.Println("\n❌ FALLO: No hay conexión o Google está caído.")
    } else {
        fmt.Println("\n✅ ÉXITO: El ping ha respondido correctamente.")
    }
}
