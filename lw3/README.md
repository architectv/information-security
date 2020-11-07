# Лабораторная работа 3 | Laboratory work 3

## Шифр DES | [Data Encryption Standard (DES)](https://en.wikipedia.org/wiki/Data_Encryption_Standard)

### Схема | Schema

![des_schema](data/in/des_schema.jpg)

### Таблицы | Tables

![des_tables](data/in/des_tables.jpg)

![des_tables2](data/in/des_tables2.jpg)

---

## Пример | Example

Используется режим "электронной кодовой книги" (простая замена).

The [**e**lectronic **c**ode **b**ook mode](https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Electronic_codebook_(ECB)) is used.

```
$ make build
go build -o des-ecb main.go
$ ./des-ecb
Usage: ./des-ecb input_file
$ ./des-ecb data/1.txt
Successfully done!
$ cat data/1.txt
qwertyui
$ cat data/1.txt.enc
<encrypted text>
$ cat data/1.txt.dec
qwertyui
```

## Дополнительные материалы | Additional materials

<img src="data/in/des1.png" alt="des1" width="400"/>

<img src="data/in/des2.png" alt="des2" width="400"/>

<img src="data/in/des3.png" alt="des3" width="400"/>
