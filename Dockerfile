FROM golang:1.12-alpine AS builder
RUN apk --update add ca-certificates git make g++
ENV GO111Module=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ARG COMMIT_HASH
ENV COMMIT_HASH=${COMMIT_HASH}
ARG BUILD_DATE
ENV BUILD_DATE=${BUILD_DATE}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -o app


FROM scratch

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/app .
COPY --from=builder /app/db ./db
EXPOSE 8080
CMD [ "./app"]