# Laboratorio 3

## Parte 1: Rompiendo ECB en imágenes

### Ejemplos

#### Tux

Con la imágen proveída de ejemplo, se pueden ver claras diferencias entre ambos cifrados. Con el cifrado AES en modo
ECB la imágen sigue siendo reconocible. Los cambios surgieron más en cuanto a los colores de la imágen, pero se mantuvieron
los detalles de la misma. Aún se puede diferenciar a Tux.

En cambio, con el cifrado AES en modo CBC la imágen es totalmente irreconocible. Lo único que hay en la imagen resultante
es ruido, no se puede distinguir a simple vista ninguna forma ni patrón.

- AES en modo ECB:

![TUX-AES-ECB](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-3/outputs/tux_aes_ecb.ppm)

- AES en modo CBC:

![TUX-AES-CBC](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-3/outputs/tux_aes_cbc.ppm)

