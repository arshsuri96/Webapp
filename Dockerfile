FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
# EXPOSE 80
# ENTRYPOINT [ "./main" ]

FROM gcr.io/distroless/base
COPY --from=builder /app/main .
EXPOSE 80
CMD ["/main"]