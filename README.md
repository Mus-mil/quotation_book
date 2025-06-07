# 📚 Quotation Book

## 🚀 Возможности

- Добавление новой цитаты
- Получение всех цитат
- Получение случайной цитаты
- Получение цитат по автору
- Удаление цитаты по ID

## 📁 Структура проекта

```
quotation_book/
├── cmd/
│   └── quotation_book/       # Точка входа в приложение
├── config/                   # Конфигурационные файлы
├── internal/
│   ├── delivery/
│   │   └── http/
│   │       └── handlers/     # HTTP-хендлеры
│   ├── models/               # Определения моделей данных
│   └── service/              # Бизнес-логика
├── mocks/                    # Сгенерированные моки для тестов
├── go.mod                    # Модуль Go
├── go.sum                    # Контрольные суммы зависимостей
├── Makefile                  # Скрипты для сборки и тестирования
└── README.md                 # Документация проекта
```
## Добавление кофигураций

файл .env  
```
PG_PASSWORD=YOUR_PASSWORD
```
Где YOUR_PASSWORD ваш пароль от postgresql  
А также добавьте конфигурации в файл config/configs.yaml

## ⚙️ Установка и запуск

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/Mus-mil/quotation_book.git
   cd quotation_book
   ```

2. **Установите зависимости:**

   ```bash
   go mod tidy
   ```

3. **Запустите приложение:**

   ```bash
   go run ./cmd/quotation_book
   ```

   или

   ```bash
   make
   ```

   Приложение будет доступно по адресу: `http://localhost:8080`

## 🧪 Тестирование

Для запуска тестов используйте команду:

```bash
go test ./...
```

Или с использованием `make`:

```bash
make test
```

## 📬 API Эндпоинты

| Метод | Путь              | Описание                            |
|-------|-------------------|-------------------------------------|
| POST  | `/quotes`         | Добавить новую цитату               |
| GET   | `/quotes`         | Получить все цитаты                 |
| GET   | `/quotes/random`  | Получить случайную цитату           |
| GET   | `/quotes?author=` | Получить цитаты по автору           |
| DELETE| `/quotes/:id`     | Удалить цитату по ID                |

## 🧰 Примеры запросов

**Добавление цитаты:**

```bash
curl -X POST http://localhost:8080/quotes \
     -H "Content-Type: application/json" \
     -d '{"author": "Лев Толстой", "quote": "Все думают о том, чтобы изменить мир, но никто не думает о том, чтобы изменить себя."}'
```

**Получение всех цитат:**

```bash
curl http://localhost:8080/quotes
```

**Получение цитат по автору:**

```bash
curl http://localhost:8080/quotes?author=Лев%20Толстой
```

**Удаление цитаты по ID:**

```bash
curl -X DELETE http://localhost:8080/quotes/1
```
