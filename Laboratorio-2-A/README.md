# Laboratorio 2 - A

## Problemas a resolver

En este laboratorio implementaremos funciones que convierten cadenas a bits y a Base 64.
Además, la investigación de una propiedad estadística de la función XOR.
1. Implementar una función para convertir una cadena de caracteres a bits. Por cada
   carácter de la cadena encontrar la representación en bytes (8 bits) del valor ASCII de
   dicho carácter. La función debe de devolver la concatenación de todos los bits de la
   cadena.

   - a) Muestre 2 ejemplos sencillos de convertir cadenas a bytes

2. Implementar una función para convertir una cadena de bytes a caracteres. Por cada
   grupo de 8 bits encontrar su representante correspondiente en ASCII. La función debe
   de devolver el texto correspondiente.

   - a) Muestre 2 ejemplos sencillos de convertir bytes a cadena

3. Implementar funciones que permitan convertir una cadena de caracteres a Base64, para
   esto utilizar la conversión manual (texto a binario, binario a código UNICODE).

   - a) Mostrar 2 ejemplos sencillos de convertir una cadena a base 64.

4. Implementar funciones que permitan convertir una cadena de base 64 a su texto
   correspondiente para esto utilizar la conversión manual (texto UNICODE a binario ,
   binario a Codigos ASCII).
   
   - a) Mostrar 2 ejemplos sencillos de convertir una cadena de base64 a su texto
      correspondiente.

5. Implementar una función que haga la operación XOR, bit a bit, con dos cadenas de
   texto.
   a. Recuerde que la llave debe ser de menor o igual tamaño que la palabra
   b. Si en dado caso la llave es menor complementarla para llegar al mismo tamaño


## Resultados

```text
Laboratorio-2-A

1. Cadenas a bytes
Hola: 01001000 01101111 01101100 01100001 
Adios: 01000001 01100100 01101001 01101111 01110011 

2. Bytes a cadenas
01001000 01101111 01101100 01100001: Hola
01000001 01100100 01101001 01101111 01110011: Adios

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