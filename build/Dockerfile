FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN make

FROM alpine

WORKDIR /build

COPY --from=builder /build/notes-server /build/notes-server

CMD [". /notes-server"]
