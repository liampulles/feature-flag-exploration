FROM golang:latest AS builder
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /go-app .

FROM scratch
COPY --from=builder /go-app .
ENTRYPOINT [ "./go-app" ]