package main

import (
    "fmt"
    "os/exec"
    "strings" // <--- NUEVO: Para poder buscar texto
)

func main() {
    target := "google.com"
    fmt.Println("--- ESPERANDO RESPUESTA DE", target, "---")

    // 1. Preparamos el comando igual que ayer
    cmd := exec.Command("ping", "-c", "1", target)

    // 2. CAMBIO IMPORTANTE: Ya no usamos cmd.Stdout = os.Stdout
    // Usamos .Output() que devuelve lo que pasó (en bytes) y si hubo error.
    outputBytes, err := cmd.Output()

    // 3. Convertimos los bytes extraños a Texto legible (String)
    resultadoTexto := string(outputBytes)

    if err != nil {
        fmt.Println("❌ Error: El comando falló o no hay internet.")
        return // Sale del programa
    }

    // 4. ANÁLISIS INTELIGENTE
    // El programa lee el texto por nosotros.
    // En Linux, un ping exitoso siempre contiene "ttl="
    if strings.Contains(resultadoTexto, "ttl=") {
        fmt.Println("✅ ÉXITO: Se encontró la palabra clave 'ttl='.")
        fmt.Println("📝 REPORTE COMPLETO RECIBIDO:")
        fmt.Println(resultadoTexto)
    } else {
        fmt.Println("⚠️ OJO: El comando corrió, pero la respuesta es rara.")
    }
} 
