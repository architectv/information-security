# Лабораторная работа 5 | Laboratory work 5

## Электронная подпись | [Digital signature](https://en.wikipedia.org/wiki/Digital_signature)

Эта программа основана на RSA.

This program is based on RSA.

## Пример | Example

```
$ make build
go build -o dsa-app main.go
$ ./dsa-app
Usage: ./dsa-app file flag
flag:
        s - sign file
        v - verify sign
$ cat data/files/text.txt
The code must be like a piece of music :)
$ ./dsa-app data/files/text.txt s
Successfully signed!
$ ./dsa-app data/files/text.txt v
Successfully verified!
$ <edit data/files/text.txt with nano or any text editor>
$ ./dsa-app data/files/text.txt v
ERROR => crypto/rsa: verification error
```
