# Ejercicio 1

## Parte 1

- Realizar la investigación de un cifrado a su elección del contenido "historia de criptografía"
- Mostrar un ejemplo de aplicación
- Explicar porque lo eligieron, cual creen que son sus ventajas y vulnerabilidades.

## Parte 2

- Realizar la creación de un script que permita la conversión de palabras en texto ASCII a BINARIO
- Realizar la creación de un script que permita la conversión de palabras en texto BASE64 a BINARIO
- Realizar la creación de un script que permita la conversión de BINARIO a BASE64
- Realizar la creación de un script que permita la conversión de BINARIO a ASCII
- Realizar la creación de un script que permita la conversión de BASE64 a ASCII
- Recuerda pasar por binario para poder hacer las conversiones
- Realizar la creación de un script que permita aplicar XOR a un BINARIO
- Realizar la creación de un script que permita generar llaves dinámicas (utilizar ASCII)
- Realizar la creación de un script que generar un nuevo cypher en ASCII con una llave k de tamaño fijo
- Realizar la creación de un script que generar un nuevo cypher en ASCII con una llave k de tamaño dinámico

### Resultados de ejecución

A continuación se adjunta el resultado de ejecutar el script `main.go`:

```shell
go run main.go
```

```text
Texto: Hello World
ASCII a binario: 01001000 01100101 01101100 01101100 01101111 00100000 01010111 01101111 01110010 01101100 01100100 
Base64 a binario: 010010000110010101010010000001000110001000000110000011000010001001011000000001000110010000
Binario a base 64: SGVsbG8gV29ybGQ
Binario a ASCII: Hello World
Base64 a ASCII: Hello world
ASCII a Base64: SGVsbG8gV29ybGQ=
XOR de 10101010 y 01010101: 11111111
Llave dinámica: #Rc$I|$v&`->EIC
Cypher fijo: =9
Cypher dinámico: ~0TZ:a:JZ1
```