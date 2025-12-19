FROM golang:1.22
WORKDIR /app
COPY . .
RUN go build -o servidor-web

# --- NUEVO: Documentamos que este contenedor usa el puerto 8080
EXPOSE 8080

CMD ["./servidor-web"]
