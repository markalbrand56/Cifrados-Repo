# Laboratorio 3

## Parte 1: Rompiendo ECB en imágenes

### Ejemplos

> [!NOTE]
> Github no fue capaz de previsualizar las imágenes cifradas, por lo que si se desea verlas
> se deberán de descargar desde la carpeta `outputs` del repositorio.

#### Tux

Con la imágen proveída de ejemplo, se pueden ver claras diferencias entre ambos cifrados. Con el cifrado AES en modo
ECB la imágen sigue siendo reconocible. Los cambios surgieron más en cuanto a los colores de la imágen, pero se mantuvieron
los detalles de la misma. Aún se puede diferenciar a Tux.

En cambio, con el cifrado AES en modo CBC la imágen es totalmente irreconocible. Lo único que hay en la imagen resultante
es ruido, no se puede distinguir a simple vista ninguna forma ni patrón.

#### Paisaje

![Paisaje original](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-3/inputs/pexels-rodolfoclix-922610.ppm)

Con esta imágen más compleja, se puede ver que el cifrado en modo ECB hace un mejor trabajo en ocultar los detalles de la imagen.
Aún se pueden ver algunos patrones, pero no es tan claro como en el caso de Tux. En el caso del cifrado en modo CBC, la imágen
nuevamente es irreconocible, siendo solo ruido en la imagen resultante.

Con una imagen más compleja, se puede ver que el cifrado en modo ECB no es tan malo como lo aparentaría el ejemplo
anterior. Aún así, el cifrado en modo CBC sigue siendo la mejor opción para cifrar imágenes.

### Preguntas para reflexión

- ¿Por qué el cifrado ECB revela los patrones de la imagen?
  - El cifrado ECB revela los patrones de la imagen porque cifra bloques de la imagen de manera independiente. Esto
    significa que si un bloque de la imagen es igual a otro, el cifrado de ambos bloques será igual. Por lo tanto, si
    hay patrones en la imagen, estos se repetirán en la imagen cifrada.

- ¿Cómo cambia la apariencia con CBC?
  - Con CBC, la apariencia de la imagen cambia drásticamente. Esto se debe a que CBC cifra los bloques de la imagen
    de manera dependiente, utilizando el cifrado de bloques anteriores para cifrar el siguiente bloque. Esto hace que
    la imagen resultante sea totalmente diferente a la original, ya que no se pueden predecir los bloques cifrados. 

- ¿Qué tan seguro es usar ECB para cifrar datos estructurados?
  - Usar ECB para cifrar datos estructurados no es seguro. Esto se debe a que, como se mencionó anteriormente, si un
    bloque de datos es igual a otro, el cifrado de ambos bloques será igual. Esto hace que sea fácil para un atacante
    identificar patrones en los datos cifrados y, por lo tanto, descifrar la información original. Aún así, como lo pudo
    demostrar el ejemplo del paisaje, no siempre es tan fácil identificar patrones en los datos cifrados con ECB.