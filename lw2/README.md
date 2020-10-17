# Лабораторная работа 2

## Алгоритм Энигмы

Программа _enigma_ с помощью алгоритма Энигмы кодирует файл _input_ и записывает результат в файл _output_.

Начальная настройка Энигмы находится в файле _settings.txt_.

Используется 3 ротора и рефлектор, алфавит содержит 256 символов.

---

# Laboratory work 2

## Enigma Algorithm

The _enigma_ program uses the Enigma algorithm to encode the _input_ file and write the result to the _output_ file.

The initial Enigma settings are located in the _settings.txt_ file.

It uses 3 rotors and a reflector, the alphabet contains 256 characters.

---

## Пример / Example

```
$ ls data
1.txt
$ cat data/1.txt
hello
$ ./enigma
Usage: ./enigma input output
$ ./enigma data/1.txt data/2.txt
Successfully done!
$ cat data/2.txt
<encrypted text>
$ ./enigma data/2.txt data/3.txt
Successfully done!
$ cat data/3.txt
hello
```