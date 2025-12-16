# 1. Variables (Para no repetir el nombre de la imagen)
IMAGE_NAME = mi-herramienta-red

# 2. Botón "build": Construye la imagen
build:
	sudo docker build -t $(IMAGE_NAME) .

# 3. Botón "run": Ejecuta con Google por defecto
run:
	sudo docker run $(IMAGE_NAME)

# 4. Botón "test": Ejecuta probando con Cloudflare
test:
	sudo docker run -e TARGET=1.1.1.1 $(IMAGE_NAME)

# 5. Botón "clean": Borra la imagen para ahorrar espacio
clean:
	sudo docker rmi $(IMAGE_NAME)
