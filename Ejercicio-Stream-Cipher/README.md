# Ejercicio Stream Cipher

## Uso

Para ejecutar el programa, se debe correr el siguiente comando:

```bash
go run main.go
```

Para ejecutar las pruebas unitarias, se debe correr el siguiente comando:

```bash
go test
```

## Desarrollo

### Generación de Keystream

Se utiliza la función `GenerateKeystream` del paquete `scripts/keys` que recibe la longitud deseada del keystream
y una semilla. Con esto, se generan números pseudoaleatorios con el algoritmo LCG implementado en el paquete
`scripts/random` de este proyecto.

```text
Keystream: [9 212 35 38 77 72 7 186 209 252 43 142 149]
```

### Cifrado

Para cifrar un mensaje, se utiliza la operación XOR entre el mensaje y el keystream. Para esto, se utiliza la función
`XORBytes` del paquete `algorithms` que recibe el mensaje y la llave.

```text
Ciphertext: [65 177 79 74 34 100 39 205 190 142 71 234 180]
```

### Descifrado

Para descifrar un mensaje, se utiliza la operación XOR entre el mensaje cifrado y el keystream. Para esto, se utiliza
la función `XORBytes` del paquete `algorithms` que recibe el mensaje cifrado y la llave.

```text
Plaintext: Hello, world!
```

### Preguntas a responder

- ¿Qué sucede cuando cambias la clave utilizada para generar el keystream?
  - Al cambiar la clave, el keystream generado cambia completamente. Esto significa que el cifrado y descifrado de un mensaje con un keystream generado con una clave diferente no es posible.
- ¿Qué riesgos de seguridad existen si reutilizas el mismo keystream para cifrar dos mensajes diferentes?
  - Si se reutiliza el mismo keystream para cifrar dos mensajes diferentes, se puede obtener información sobre los mensajes originales. Esto se debe a que al realizar la operación XOR entre el keystream y los mensajes, se pueden obtener los mensajes originales.
- ¿Cómo afecta la longitud del keystream a la seguridad del cifrado?
  - La longitud del keystream afecta la seguridad del cifrado, ya que si el keystream es muy corto, se puede realizar un ataque de fuerza bruta para obtener la clave y descifrar el mensaje. Por otro lado, si el keystream es muy largo, se puede realizar un ataque de frecuencia para obtener la clave y descifrar el mensaje. 
- ¿Qué consideraciones debes tener al generar un keystream en un entorno real?
  - Al generar un keystream en un entorno real, se debe tener en cuenta la semilla utilizada para generar el keystream, ya que si la semilla es predecible, se puede obtener el keystream y descifrar el mensaje. Además, se debe tener en cuenta la longitud del keystream, ya que si es muy corto, se puede realizar un ataque de fuerza bruta para obtener la clave y descifrar el mensaje.

## Uso de IA generativa

Se adjuntan los recursos de IA generativa utilizados para la creación de este ejercicio:

- [Chat-GPT (Free Tier)](https://chatgpt.com/share/67bbc25e-5668-8012-9300-c2d5acc0c67b)