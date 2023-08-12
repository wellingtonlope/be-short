FROM golang:1.21-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o main ./cmd

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]