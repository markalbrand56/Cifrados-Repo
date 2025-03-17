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

## Parte 2: Capturando Cifrado en Red con Wireshark

### Preguntas para reflexión:

- ¿Se puede identificar que los mensajes están cifrados con AES-CBC?

- ¿Cómo podríamos proteger más esta comunicación?

## Parte 3: Implementando un Cifrado de Flujo con ChaCha20

Para medir el rendimiento del cifrado de flujo con ChaCha20 en comparación con AES-CBC, se utilizó el profiler
de Go. Se utilizó una imagen de 33.2 MB para realizar las pruebas.

```text
Parte 3: Implementando un Cifrado de Flujo con ChaCha20
Cantidad de bytes: 34682396
AES-CBC → Tiempo: 94.8538ms, Memoria usada: 110091 KB
ChaCha20 → Tiempo: 83.7184ms, Memoria usada: 33869 KB
```

###  Preguntas para reflexión:

- ¿Analizar que cifrado es mas rápido ChaCha20 o AES?
  - En este caso, ChaCha20 resultó ser más rápido que AES-CBC. Aunque la diferencia no es muy grande, ChaCha20
    logró cifrar la imagen en menos tiempo que AES-CBC. En cuanto a la memoria utilizada, ChaCha20 también utilizó
    menos memoria que AES-CBC, alrededor de 76 MB menos lo cuál es una diferencia considerable.

- ¿En qué casos debería usarse en vez de AES?
  - ChaCha20 debería usarse en casos donde se requiera un cifrado rápido y seguro. Aunque AES es un algoritmo
    de cifrado muy seguro, ChaCha20 es más rápido y eficiente en términos de rendimiento. Por lo tanto, si se
    necesita cifrar grandes cantidades de datos de manera rápida, ChaCha20 es una buena opción. 