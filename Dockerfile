FROM golang:latest AS builder

COPY . /app

WORKDIR /app

# Download Go modules
RUN go get task2
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/task2 cmd/main.go

FROM scratch

ENV PORT=8080

LABEL maintainer="Sergey Nesterenko <Sergey@Nesterenko.net>"
LABEL version=0.1

WORKDIR /app
COPY --from=builder /build/task2 task2
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE ${PORT}

# Run
CMD ["/app/task2"]