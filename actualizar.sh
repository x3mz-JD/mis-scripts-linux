#!/bin/bash

echo "Hola, iniciando actualización del sistema..."
sudo apt update && sudo apt upgrade -y
echo "¡Sistema actualizado correctamente!"
neofetch
