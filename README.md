#  Marketplace

Простое RESTful API для маркетплейса, написанное на **Go** с использованием **Gin**, **GORM**, и **PostgreSQL**.

---

##  Функционал

- Регистрация пользователей  
- Авторизация (JWT)  
- Создание объявлений (ads)  
- Получение списка объявлений

---

## ⚙️ Установка и запуск

### 1. Клонируй репозиторий

```bash
git clone https://github.com/yourusername/marketplace.git
cd marketplace
```
### 2. Настрой переменные окружения (если запускаешь без Docker)

```bash
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=marketplace
export DB_PORT=5432
export JWT_SECRET=your_jwt_secret
```
### 3. Создай базу данных PostgreSQL

Подключись к Postgres:

```bash
psql -U postgres
```
```bash
CREATE DATABASE marketplace;
\q
```
### 4. Запуск локально
```bash
   go run cmd/main.go
```
### 5. Запуск через Docker Compose

```bash
docker compose up --build
```
Это поднимет:
- контейнер с Postgres (порт 5432)
- приложение Go (порт 8080)

