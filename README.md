## Для ручного запуска:
1. `go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0`
2. `oapi-codegen -generate gorilla,types wallet/api/frontend/frontendapi.yaml > wallet/api/frontend/frontendapi.gen.go`
3. `go mod tidy -v`
4. `go run main.go`

## Для работы с БД
1. Установить PostgreSQL
2. Миграции выполняются автоматически

## Миграции
up и down миграции выполняются автоматически при запуске приложения

[Инструкция по установке утилиты для миграций](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

Создание миграции: `migrate create -ext sql -dir {path-to-migrations-dir} {migartion-name}`

## Основные маршруты API

- **POST /banktask/api/v1/wallet** - Изменение баланса кошелька. В теле запроса передаётся uuid кошелька, тип операции (DEPOSIT или WITHDRAW) и число, на которое изменяется баланс.
- **GET /banktask/api/v1/wallets/{wallet_id}** - Получение баланса кошелька. В пути передаётся uuid кошелька, баланс которого нужно получить.

## Зависимости

Проект использует следующие библиотеки:

- **Gorilla Mux**: HTTP-маршрутизатор для Go.
- **pgx**: Драйвер для работы с базой данных PostgreSQL.
- **oapi**: Инструмент для генерации Go-кода на основе спецификаций OpenAPI.
- **godotenv**: Для загрузки переменных окружения из файла .env.
