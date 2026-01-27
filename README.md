# Wallet App

Небольшой сервис кошельков на Go. Позволяет создавать кошельки, получать баланс и изменять его (пополнение/списание) с сохранением данных в PostgreSQL. Проект упакован в Docker и запускается через `docker-compose`.

---

## Что здесь происходит

Проект состоит из нескольких частей:

* **HTTP‑сервис** - основной API (Gin)
* **PostgreSQL**
* **Миграции** — SQL‑скрипты для создания и обновления схемы БД
* **Docker / docker‑compose**

Структура проекта (упрощённо):

```
wallet-app/
├──bruno/               # схема для запуска курлов в bruno
├── cmd/
│   ├── service/        # запуск сервиса
│   │   └── main.go
│   └── migrate/        # запуск миграций
│       └── main.go
├── internal/           # доменная логика, репозитории, сервисы, хэндлеры
├── migrations/
│   └── changelog/
│       └── master/     # .sql файлы миграций
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## API (кратко)

Примеры того, что умеет сервис (ручки):

* Создать кошелёк
* Получить баланс по `wallet_id`
* Изменить баланс (delta‑операция)

---

## Требования

Перед запуском убедись, что установлено:

* **Docker** ≥ 20
* **Docker Compose** (v2)

---

## Установка и запуск

### 1. Клонировать репозиторий

```bash
git clone https://github.com/ChurchHEllla/wallet-app.git
```

ИЛИ
Скачайте [ZIP](https://github.com/ChurchHEllla/wallet-app/archive/refs/heads/main.zip) проекта

### 2. Собрать и запустить проект через Docker

#### docker-compose
```bash
docker compose-up --build
ИЛИ
make docker-up
```
#### Миграция на docker
```bash
make docker-migrate
```
#### Запуск тестов хэндлеров
```bash
make test-handler
```
---

## Переменные окружения

Они задаются в `.env`:
* `POSTGRES_DB=wallet` - имя базы
* `POSTGRES_PORT=5432` - порт подключения
* `POSTGRES_USER=wallet_user` - пользователь
* `POSTGRES_PASSWORD=wallet_pass` - пароль
* `DATABASE_HOST=db` - строка подключения (используется сервисом и миграциями)
