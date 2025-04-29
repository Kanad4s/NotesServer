FROM golang:alpine AS builder

WORKDIR /NotesServer
RUN apk add --no-cache make git

COPY go.mod go.sum ./
RUN go mod download

#ADD go.mod .

COPY . .
RUN make -C ./build

FROM alpine:3.18

WORKDIR /app

# Копирование бинарника
COPY --from=builder /NotesServer/build/notes-server .

# Дополнительные файлы (конфиги, статика)
# COPY --from=builder /NotesServer/configs /etc/myapp/

CMD ["./notes-server"]
