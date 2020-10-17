# Лабораторная работа 1

## Лицензия для программы

Программа _app_ запускается корректно лишь после запуска установщика _install_, который привязывает данное приложение к компьютеру. В качестве ключа здесь используется _machine-id_.

Критерии выбора ключа:

1. Уникальность
2. Неизменяемость
3. Доступность

---

# Laboratory work 1

## License for the program

The _app_ program starts correctly only after the _install_ installer (it binds this application to the computer) has been launched. _Machine-id_ is the key.

Key selection criteria:

1. Uniqueness
2. Immutability
3. Availability

---

## Пример / Example

```
$ ./app
ERROR! There is no license!
$ ./install
Key was written successfully!
$ ./app
It's OK! You got a license for this program.
```