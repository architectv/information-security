# Лабораторная работа 4 | Laboratory work 4

## Шифр RSA | [RSA Cipher](https://en.wikipedia.org/wiki/RSA_(cryptosystem))

## Алгоритм | Algorithm

### Генерация ключей | Generating keys

```
P, Q [two different prime numbers]
N = P*Q [alphabet length]
Phi = (P - 1)*(Q - 1) [Euler function of N]
E - prime number that GCD(E, Phi) = 1
D that (E*D) mod Phi = 1
{E, N} - public key
{D, N} - private key
```

### Шифрование | Encryption

```
M - original message
C - encrypted message
C = M**E mod N
```

### Расшифрование | Decryption

```
M = C**D mod N
```

## Пример | Example

Используется режим "электронной кодовой книги" (простая замена).

The [**e**lectronic **c**ode **b**ook mode](https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Electronic_codebook_(ECB)) is used.

P, Q генерируются размером *bits* бит | P, Q are generated in size *bits* bits.

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
