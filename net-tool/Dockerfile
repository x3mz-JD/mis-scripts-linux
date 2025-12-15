# 1. BASE: Empezamos con un Linux que ya tiene Go instalado (la imagen oficial)
FROM golang:1.22

# 2. CARPETA: Creamos una carpeta de trabajo dentro del contenedor
WORKDIR /app

# --- LÍNEA NUEVA ---
# Instalamos 'iputils-ping' que es el paquete que contiene el comando ping en Debian
RUN apt-get update && apt-get install -y iputils-ping
# -------------------

# 3. COPIA: Copiamos tus archivos (main.go, go.mod) desde TU PC al CONTENEDOR
COPY . .

# 4. CONSTRUCCIÓN: Compilamos el programa dentro del contenedor
RUN go build -o net-tool

# 5. ARRANQUE: Le decimos qué comando ejecutar cuando el contenedor se inicie
CMD ["./net-tool"]
