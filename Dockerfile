FROM golang:1.24.3
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]
