FROM golang:alpine AS builder

WORKDIR /NotesServer

COPY . .

RUN apt-get update && \
    apt-get install -y \
    make \
    && \
    make -C ./build
RUN ls

FROM ubuntu:24.04

WORKDIR /build

COPY --from=builder /build/notes-server .

RUN ls
RUN pwd

CMD [". /notes-server"]
