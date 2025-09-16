# Wildberries API Integration

Сервис для интеграции с API Wildberries, предоставляющий удобный интерфейс для работы с различными эндпоинтами маркетплейса.

## Возможности

- Получение списка карточек товаров
- Управление остатками товаров
- Работа с заказами
- Получение аналитики и статистики продаж
- Обновление цен

## Требования

- Go 1.24 или выше
- Docker (опционально)

## Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/obumax/wb-api-integration.git
cd wb-api-integration
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте файл с переменными окружения:
```bash
cp .env.example .env
```

4. Заполните переменные окружения в файле `.env`:
```bash
WB_STANDARD_TOKEN=your_standard_token_here
WB_ADV_TOKEN=your_adv_token_here
WB_STAT_TOKEN=your_stat_token_here
LOG_LEVEL=info
PORT=8080
```

## Запуск

### Локальный запуск

```bash
go run cmd/main.go
```

### Через Docker

1. Соберите образ:
```bash
docker build -t wb-api-integration .
```

2. Запустите контейнер:
```bash
docker run -p 8080:8080 --env-file .env wb-api-integration
```

## API Endpoints

### Товары
- `GET /api/products` - получение списка товаров
- `POST /api/stocks` - обновление остатков
- `POST /api/prices` - обновление цен

### Заказы
- `GET /api/orders` - получение списка заказов

### Аналитика
- `GET /api/analytics` - получение аналитических данных

## Разработка

### Добавление нового эндпоинта

1. Добавьте новый метод в `internal/client/wb.go`
2. Создайте обработчик в `internal/handlers/wb_handler.go`
3. Добавьте новый роут в `cmd/main.go`

### Тестирование

```bash
go test ./...
```

## Лицензия

MIT

## Автор

Maxim Obukhov