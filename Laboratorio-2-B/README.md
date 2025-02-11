# Laboratorio 2 - A

## Problemas a resolver

Requerimiento:
1. Implementar una función que haga la operación XOR, bit a bit, con dos cadenas de
   texto.
   Problemas a resolver
2. Dada la imagen XOR_Imagen, y la llave “cifrados_2025” encontrar el valor original de la
   imagen.
   a. Deben de convertir la imagen a base 64 y aplicarle un xor con la llave para
   encontrar su valor
3. Investigar porque al aplicar XOR con una llave de texto la imagén se corrompe.
4. Investigar como aplicar un xor a 2 imagnes. Para esto deben de eleccionar 2 imágenes,
   luego proceder hacer un xor entre las dos imágenes. Esto significa que una imagen es la
   original y la otra se utilizará como llave para aplicar el xor.
   a. Mostrar las imágenes utilizadas y el resultado, asi mismo explique que
   inconvenientes encontro al momento de realizar el xor.
5. 


## Resultados

```text
Laboratorio-2-A

1. Cadenas a bytes
Hola: [72 111 108 97]
Adios: [65 100 105 111 115]

2. Bytes a cadenas
[72 111 108 97]: Hola
[65 100 105 111 115]: Adios

3. Texto a base64
Hola: SG9sYQ==
Adios: QWRpb3M=

4. Base64 a texto
SG9sYQ==: Hola
QWRpb3M=: Adios

5. XOR entre dos cadenas
01001000 01101111 01101100 01100001  (Hola) XOR 01000001 01100100 01101001 01101111 01110011  (Adios): 00001001000010110000010100001110
```


## Uso de IA generativa