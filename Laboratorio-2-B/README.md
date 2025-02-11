# Laboratorio 2 - A

## Problemas a resolver

Requerimiento:
1. Implementar una función que haga la operación XOR, bit a bit, con dos cadenas de
   texto.
2. Dada la imagen XOR_Imagen, y la llave “cifrados_2025” encontrar el valor original de la
   imagen.
   - a) Deben de convertir la imagen a base 64 y aplicarle un xor con la llave para
   encontrar su valor
3. Investigar porque al aplicar XOR con una llave de texto la imagén se corrompe.
4. Investigar como aplicar un xor a 2 imágenes. Para esto deben de seleccionar 2 imágenes,
   luego proceder hacer un xor entre las dos imágenes. Esto significa que una imagen es la
   original y la otra se utilizará como llave para aplicar el xor.
   - a) Mostrar las imágenes utilizadas y el resultado, asi mismo explique que
   inconvenientes encontró al momento de realizar el xor.


## Resultados

```bash
go run main.go
```

### XOR entre dos cadenas

```text
1. Cifrado XOR con cadenas de texto
	Hola XOR abc1 = 290d0f50
```

### Descifrado de imagen con llave

![XOR_Imagen](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-2-B/examples/imagen_descifrada.png)

### XOR entre dos imágenes

![Imagen 1](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-2-B/examples/imagen3.png)
![Imagen 2](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-2-B/examples/imagen4.png)

![Imagen XOR](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-2-B/examples/imagen_xor_resultado.png)

## Uso de IA generativa

- [ChatGPT (Free Tier)](https://chatgpt.com/share/67ab6de3-349c-8012-8079-b0fa07a84283)