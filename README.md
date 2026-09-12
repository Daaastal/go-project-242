### Hexlet tests and linter status:
[![Actions Status](https://github.com/Daaastal/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Daaastal/go-project-242/actions)

### My tests status:
[![Go](https://github.com/Daaastal/go-project-242/actions/workflows/go.yml/badge.svg)](https://github.com/Daaastal/go-project-242/actions/workflows/go.yml)

https://asciinema.org/connect/0595ed4b-3585-4766-b709-50a290dbc6ce

## Сборка
```bash
make build
```

## Тестирование
```bash
make test
```

## Примеры
```bash
./bin/hexlet-path-size -h
NAME:
   hexlet-path-size - print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)

USAGE:
   hexlet-path-size [global options] <path>

GLOBAL OPTIONS:
   --recursive, -r  recursive size of directories (default: false)
   --human, -H      human-readable sizes (auto-select unit) (default: false)
   --all, -a        include hidden files and directories (default: false)
   --help, -h       show help
```

Обычный запуск — размер одного файла:

```console
$ ./bin/hexlet-path-size hello_file.txt
7B      hello_file.txt
```

Размер каталога без рекурсии:

```console
$ ./bin/hexlet-path-size dir
20B     dir
```

С рекурсией, обход вложенных каталогов:

```console
$ ./bin/hexlet-path-size -r dir
503B    dir
```

С флагом `-a` - учитывает скрытые файлы и каталоги внутри:

```console
$ ./bin/hexlet-path-size -r -a .
```

Человекочитаемый вывод:

```console
$ ./bin/hexlet-path-size -H big_file.txt
87.9KB  big_file.txt
```

Скрытый файл, переданный как корневой путь, измеряется всегда:

```console
$ ./bin/hexlet-path-size .hidden
6B      .hidden
```

## Поведение ошибок

Пустой запуск и лишний аргумент дают одинаковое сообщение и код выхода 2:

```console
$ ./bin/hexlet-path-size
usage: hexlet-path-size [flags] <path>
$ echo $?
2
```

Несуществующий путь даёт код 1 и сообщение об ошибке файловой системы:

```console
$ ./bin/hexlet-path-size no-such-file
lstat "no-such-file": no such file or directory
$ echo $?
1
```
