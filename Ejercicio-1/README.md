# Ejercicio 1

## Parte 1

- Realizar la investigación de un cifrado a su elección del contenido "historia de criptografía"
- Mostrar un ejemplo de aplicación
- Explicar porque lo eligieron, cual creen que son sus ventajas y vulnerabilidades.

### Historia de la criptografía: Cifrado Atbash

El cifrado Atbash es uno de los cifrados más antiguos conocidos en la actualidad, datando de alrededor del 500 a.C. 
Originalmente fue utilizado por los hebreos para cifrar mensajes secretos, incluso lo utilizaron para ocultar mensajes
en la biblia.

A lo largo del tiempo este cifrado ha sido utilizado por más culturas y civilizaciones, como los griegos y los romanos.

Este cifrado es un tipo de cifrado de sustitución en el que las letras del alfabeto se sustituyen por sus equivalentes inversos. 
Por ejemplo, la primera letra del alfabeto, la "A", se sustituye por la última letra del alfabeto, la "Z". 

El cifrado Atbash es un tipo de cifrado monoalfabético, es decir, cada letra del alfabeto se sustituye por una única letra.

#### Referencias

- [Understanding the Atbash Cipher: A Simple and Effective Encryption Method | 2023](https://cyberw1ng.medium.com/understanding-the-atbash-cipher-a-simple-and-effective-encryption-method-2023-da89d71c4369)

### Ejemplo de aplicación

En el archivo `algorithms/atbash.go` se encuentra la siguiente implementación:

```go
package algorithms

// Atbash Cifra un texto utilizando el cifrado Atbash.
func Atbash(s string) string {
	var result string
    for _, c := range s {
        // Si el carácter es una letra minúscula, se le resta a 'z' y se suma 'a'
        if c >= 'a' && c <= 'z' {
            result += string('z' - c + 'a')
        }
        // Si el carácter es una letra mayúscula, se le resta a 'Z' y se suma 'A'
        if c >= 'A' && c <= 'Z' {
            result += string('Z' - c + 'A')
        }
    }
	return result
}
```

En esta aplicación, se le pasa un texto a la función `Atbash` y esta retorna el texto cifrado utilizando el cifrado Atbash.

Empieza por un bucle, iterando sobre cada carácter del texto. Si el carácter es una letra minúscula, se le resta a 'z' y se suma 'a'.
Si el carácter es una letra mayúscula, se le resta a 'Z' y se suma 'A'.

Esto resulta en la letra inversa del alfabeto, es decir, la 'A' se convierte en 'Z', la 'B' en 'Y', y así sucesivamente.

Por ejemplo, si la letra es 'b' se convierte en 'y', si la letra es 'A' se convierte en 'Z'.

Con esta implementación, una aplicación práctica sería cifrar un mensaje secreto para enviarlo a otra persona de forma segura.

```go
package main

import (
    "fmt"
    "Ejercicio-1/algorithms"
    "Ejercicio-1/utils"
)

func main() {
    text := "Hello World"
    cipher := algorithms.Atbash(text)
    
    err := sendSecretMessage(cipher, "John Doe")
	
    if err != nil {
        fmt.Println(err)
    }
}
```

### Ventajas y vulnerabilidades

#### Ventajas

- Es un cifrado muy simple y fácil de entender.
- Es un cifrado simétrico, es decir, el cifrado y descifrado son iguales.

#### Vulnerabilidades

- Es un cifrado muy simple y fácil de romper, ya que solo se necesita conocer el alfabeto para descifrar el mensaje.
- Es un cifrado monoalfabético, es decir, cada letra del alfabeto se sustituye por una única letra, lo que lo hace más vulnerable.
- Es propenso a ataques de fuerza bruta, ya que solo hay 26 posibles combinaciones.

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
Llave dinámica: P:l[&r}0+\@0>7Q
Cypher fijo: PPP
Descifrado (fijo): abc
Cypher dinámico: ! SCV,9O
Descifrado (dinámico): Hello World
```