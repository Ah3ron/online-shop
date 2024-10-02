FROM node:22 AS frontend-build
WORKDIR /app/frontend
RUN npm install -g pnpm
COPY frontend/package*.json ./
RUN pnpm install
COPY frontend/ .
RUN pnpm run build
RUN mv build views

FROM golang:1.23.1 AS backend-build
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod tidy
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=frontend-build /app/frontend/views ./views
COPY --from=backend-build /app/backend/server ./
EXPOSE 3000
CMD ["./server"]

