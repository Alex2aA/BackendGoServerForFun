# BackendGoServerForFun

Два Go-микросервиса (user-service, booking-service) с PostgreSQL для управления пользователями, хостелами, домами и бронированиями.

## Архитектура

```
                    ┌─────────────┐
                    │  Frontend   │
                    │ (Vue 3 SPA) │
                    └──────┬──────┘
                           │
              ┌────────────┼────────────┐
              │ 8081       │ 8082       │
              ▼            ▼            │
      ┌────────────┐ ┌────────────┐     │
      │ user-      │ │ booking-   │     │
      │ service    │ │ service    │     │
      │ (Go 1.24)  │ │ (Go 1.24)  │     │
      └─────┬──────┘ └─────┬──────┘     │
            │              │            │
            └──────────────┼────────────┘
                           │
                           ▼ 5432
                    ┌────────────┐
                    │ PostgreSQL │
                    │    16      │
                    └────────────┘
```

### user-service (порт 8081)
- Регистрация / логин / логаут
- JWT-аутентификация (access + refresh token)
- Clean architecture: `delivery/http` → `usecase` → `repository/postgres`

### booking-service (порт 8082)
- CRUD хостелов, домов, бронирований
- JWT-аутентификация (общая с user-service)
- Миграции БД (создают таблицы `users`, `hostels`, `free_houses`, `booked`)

## Быстрый старт

```bash
# 1. Запустить все сервисы
docker compose up --build -d

# 2. Проверить логи
docker compose logs -f

# 3. Остановить
docker compose down
```

После запуска:
- `POST http://localhost:8081/api/register` — регистрация
- `POST http://localhost:8081/api/login` — логин
- `POST http://localhost:8082/api/hostel` — создание хостела
- `POST http://localhost:8082/api/house` — создание дома
- `POST http://localhost:8082/api/booking` — бронирование

## Структура проекта

```
BackendGoServerForFun/
├── docker-compose.yaml          # Оркестрация 3 контейнеров
├── test-requests.sh             # Набор curl-запросов для тестирования
├── BackendGoServerForFun.postman_collection.json
├── user-service/                # Микросервис пользователей
│   ├── cmd/app/main.go          # Точка входа
│   ├── config/config.go         # Конфигурация
│   ├── internal/
│   │   ├── delivery/http/       # HTTP-хендлеры
│   │   ├── domain/              # Сущности (User, Token)
│   │   ├── middleware/          # JWT-мидлварь
│   │   ├── repository/          # Репозиторий (PostgreSQL)
│   │   └── usecase/             # Бизнес-логика
│   ├── pkg/logger/              # Логгер (zap)
│   └── Dockerfile
├── booking-service/             # Микросервис бронирований
│   ├── cmd/app/main.go          # Точка входа
│   ├── config/config.go         # Конфигурация
│   ├── migrations/              # SQL-миграции
│   │   ├── 000001_create_users.up.sql
│   │   ├── 000002_create_hostels_houses.up.sql
│   │   └── 000003_create_booked.up.sql
│   ├── internal/
│   │   ├── delivery/http/       # Хендлеры (hostel, house, booking)
│   │   ├── domain/              # Сущности
│   │   ├── middleware/          # JWT-мидлварь
│   │   ├── repository/          # Репозитории (PostgreSQL)
│   │   └── usecase/             # Бизнес-логика
│   ├── pkg/logger/              # Логгер (zap)
│   └── Dockerfile
```

## Фронтенд

Для удобного тестирования API можно использовать Vue-приложение:

```
https://github.com/Alex2aA/BackendGoServerForFun-frontend
```

Фронтенд автоматически проксирует запросы на user-service (порт 8081) и booking-service (порт 8082).

## API Endpoints

### user-service

| Method | Path              | Описание       |
|--------|-------------------|----------------|
| POST   | `/api/register`   | Регистрация    |
| POST   | `/api/login`      | Вход           |
| POST   | `/api/logout`     | Выход          |
| GET    | `/api/user`       | Инфо о юзере   |

### booking-service

| Method | Path              | Описание          |
|--------|-------------------|-------------------|
| POST   | `/api/hostel`     | Создать хостел    |
| GET    | `/api/hostel/:id` | Получить хостел   |
| POST   | `/api/house`      | Создать дом       |
| GET    | `/api/house/:id`  | Получить дом      |
| POST   | `/api/booking`    | Создать бронь     |
