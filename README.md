## To-do-app 

## Назначение сервиса
Сервис предназначен для управления списком задач (To-Do List)

## Владелец/ответственный за сервис
Ответственность за сервис несет команда [Алуа](https://github.com/aluamakhanova).

## Системные требования
Для работы сервиса необходимы:
- go
- база данных MySql

## Переменные окружения
Список переменных окружения содержится в файле `.env.example`.

## Запуск проекта локально
  ```bash
    go run main.go
  ```
## Запуск миграций
  ```bash
dbmate up
  ```

## Инициализация нового проекта
  ```bash
    go mod tidy
    go build -o main .
  ```

## Структура проекта
    - main.go - точка входа приложения
    - config/config.go - инициализация конфигурации приложения
    - db/migrations - файлы миграции
    - internal/tasks - основная часть приложения
    - internal/test_tasks - тесты