# Laboratorio 3

## Parte 4

### Compilación

```bash
GOOS=linux GOARCH=amd64 go build -o build/ransomware encryptor.go
```

```bash
GOOS=linux GOARCH=amd64 go build -o build/decryptor decryptor.go
```

### Demostración

En el ejemplo, se utilizó una máquina virtual con Kali Linux, donde se ejecutó el ransomware y posteriormente el decryptor.

En la captura de pantalla se puede observar como el ransomware cifra los archivos de la carpeta en la que se encuentra, y posteriormente el decryptor descifra los archivos.

![Captura de pantalla]()
