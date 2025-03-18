# Laboratorio 3

## Parte 2

![Wireshark](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-3-Parte-2/images/wireshark.png)

### Preguntas para reflexión:

- ¿Se puede identificar que los mensajes están cifrados con AES-CBC?
  - No del todo. En la sección de data del paquete interceptado se puede ver lo que el cliente le envió al servidor, pero lo que se intercepta son solo bytes sin sentido. No se puede identificar que los mensajes están cifrados con AES-CBC de una manera tan directa. Se podría analizar los bytes en una sucesión de comunicaciones para intentar analizar los vectores IV y las claves de cifrado, pero no es algo trivial.

- ¿Cómo podríamos proteger más esta comunicación?
  - Un primer paso sería utilizar una conexión segura. En esta prueba local se utiliza HTTP, lo que significa que los mensajes viajan en texto plano. Si se utilizara HTTPS, los mensajes viajarían cifrados y sería mucho más difícil interceptarlos. Además, se podría utilizar un cifrado más seguro que AES-CBC, como AES-GCM, que es más seguro y eficiente.