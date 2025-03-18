# Laboratorio 3

## Parte 4

### Compilación

Para compilar el ransomware y el decryptor, se debe ejecutar el siguiente comando en la terminal.
Esto generará un archivo ejecutable en la carpeta build para ejecutarse en sistemas Linux.

> [!WARNING]
> Es importante tener en cuenta que el ransomware cifrará los archivos de la carpeta en la que se encuentre, por lo que se recomienda probarlo en un entorno controlado.

```bash
GOOS=linux GOARCH=amd64 go build -o build/ransomware encryptor.go
```

```bash
GOOS=linux GOARCH=amd64 go build -o build/decryptor decryptor.go
```

### Demostración

En el ejemplo, se utilizó una máquina virtual con Kali Linux, donde se ejecutó el ransomware y posteriormente el decryptor.

En la captura de pantalla se puede observar como el ransomware cifra los archivos de la carpeta en la que se encuentra, y posteriormente el decryptor descifra los archivos.

![Captura de pantalla](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-3-Parte-4/images/ransomware.png)

![Captura de pantalla](https://github.com/markalbrand56/Cifrados-Repo/blob/main/Laboratorio-3-Parte-4/images/decryptor.png)

### Preguntas para reflexión:

- ¿Cómo podríamos evitar ataques de ransomware?
  - Para evitar estos ataques es importante, por un lado, cuidar la seguridad de las máquinas siendo cuidadoso de los archivos que se descargan y ejecutan, y por otro lado, tener copias de seguridad de los archivos importantes.
    de esta manera, en caso de ser atacado por un ransomware, se puede recuperar la información sin tener que pagar el rescate.

- ¿Qué tan importante es almacenar claves de manera segura?
  - Es de vital importancia almacenar las claves de manera segura, ya que estas son la base de la seguridad de los sistemas. Si las claves son robadas, un atacante podría tener acceso a información sensible y comprometer la seguridad de la organización.