FROM golang:alpine AS builder

WORKDIR /NotesServer
RUN apk add --no-cache make git

RUN ls
RUN pwd

#COPY ../go.mod ../go.sum ./
#RUN go mod download

COPY . .
RUN go build -o ./notes-server ../cmd 

RUN ls
RUN pwd
RUN cat Makefile

FROM alpine:3.18

WORKDIR /app

# Копирование бинарника
COPY --from=builder /NotesServer/notes-server .

# Дополнительные файлы (конфиги, статика)
# COPY --from=builder /NotesServer/configs /etc/myapp/

RUN ls
RUN pwd

CMD ["./notes-server"]
