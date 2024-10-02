# Этап сборки SvelteKit
FROM node:18 AS frontend-build

# Установим рабочую директорию
WORKDIR /app/frontend

# Устанавливаем pnpm
RUN npm install -g pnpm

# Копируем package.json и pnpm-lock.yaml
COPY frontend/package*.json ./

# Устанавливаем зависимости с помощью pnpm
RUN pnpm install

# Копируем остальные файлы проекта
COPY frontend/ .

# Собираем проект
RUN pnpm run build

# Переименовываем папку build в views
RUN mv build views

# Этап сборки Go приложения
FROM golang:1.23.1 AS backend-build

# Установим рабочую директорию
WORKDIR /app/backend

# Копируем go.mod и go.sum
COPY backend/go.mod backend/go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем остальные файлы приложения
COPY backend/ .

# Собираем Go приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Финальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранный SvelteKit проект
COPY --from=frontend-build /app/frontend/views ./views

# Копируем собранное Go приложение
COPY --from=backend-build /app/backend/myapp ./backend

# Открываем порт 3000
EXPOSE 3000

# Указываем команду для запуска приложения
CMD ["./backend/myapp"]

