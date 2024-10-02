# Тестовое задание в Effective mobile

## Деплой
## Шаг 1: клонирование репозитория
Сначала клонируйте репозиторий на свой локальный компьютер или сервер:

``` bash
git clone https://github.com/0Abracadaber0/song-library.git
cd song-library
```

## Шаг 2: создание ```.env``` файла
В корневом каталоге вашего проекта создайте файл ```.env```, чтобы сохранить переменные окружения. Мой ```.env``` для тестирования выглядил так (файл без значений ```.env.example``` лежит в корне репозитория):
```
POSTGRES_PORT=5432
POSTGRES_HOST=postgres
POSTGRES_DB=library
POSTGRES_USER=library-user
POSTGRES_PASSWORD=12345
EXTERNAL_HOST=api_service
EXTERNAL_PORT=3001
APP_HOST=0.0.0.0
APP_PORT=8080
```
Эти значение устанвливаются по умолчанию, в случае их отсутствия в вашем окружении (кроме пароля).

## Шаг 3: настройка ```docker-compose.yaml```
Ваш файл должен выглядить следующим образом (заполненный пример находится в корне проекта):
```yaml
version: '3'
services:
  postgres:
    image: postgres:latest
    env_file:
      - .env
    ports:
      - "[выбранный вами порт]:[выбранный вами порт]"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER} -d ${POSTGRES_DB}"]
      interval: 10s
      timeout: 9s
      retries: 5
      start_period: 10s

  library:
    build:
      context: .
      dockerfile: Dockerfile
    env_file:
      - .env
    ports:
      - "8080:8080"
    depends_on:
      postgres:
        condition: service_healthy

  api_service:
    image: [образ внешнего api]
    ports:
      - "[порт внешнего api]:[порт внешнего api]"

volumes:
  postgres_data:
```
healthcheck необходим для проверки готовности контейнера postgres к подключению приложния.

## Шаг 4: сборка и запуск приложения
Чтобы собрать и запустить приложение, выполните следующую команду в терминале:
```
docker-compose up -d --build
```
Уберите флаг ```-d```, если хотите видеть отоброжение логов.

## P.S.
Впервые пишу инструкцию по деплою, надеюсь, что не намудрил)
