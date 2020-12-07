# Лабораторная работа 6 | Laboratory work 6

## LZW-сжатие | [LZW compression](https://en.wikipedia.org/wiki/Lempel–Ziv–Welch)

[Good source](https://www.geeksforgeeks.org/lzw-lempel-ziv-welch-compression-technique/)

## Пример | Example

```
$ make build
go build -o lzw-app main.go
$ ./dsa-app
Usage: ./lzw-app file flag
flag:
        s - compress file
        v - decompress file
$ cat data/files/test.txt
WYS*WYGWYS*WYSWYSG
$ ./lzw-app data/files/test.txt c
Successfully compressed!
$ ./lzw-app data/comp/test.txt d
Successfully decompressed!
```
