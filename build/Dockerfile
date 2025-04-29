FROM golang:alpine AS builder

WORKDIR /NotesServer

COPY . .

RUN ls
RUN pwd

RUN make -C ./build

RUN ls

FROM ubuntu:24.04

WORKDIR /app

COPY --from=builder /NotesServer/notes-server .

RUN ls
RUN pwd

CMD ["./notes-server"]
