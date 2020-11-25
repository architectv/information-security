# Лабораторная работа 4 | Laboratory work 4

## Шифр RSA | [RSA Cipher](https://en.wikipedia.org/wiki/RSA_(cryptosystem))

## Пример | Example

Используется режим "электронной кодовой книги" (простая замена).

The [**e**lectronic **c**ode **b**ook mode](https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Electronic_codebook_(ECB)) is used.

```
$ make build
go build -o rsa-ecb main.go
$ ./rsa-ecb
Usage: ./rsa-ecb input_file bits
$ cat data/in/text.txt
qwertyui
$ ./rsa-ecb data/in/text.txt 100
Encode time: 2.9992ms
Decode time: 2.9973ms
Successfully done!
$ cat data/enc/text.txt
<encrypted text>
$ cat data/dec/text.txt
qwertyui
```
