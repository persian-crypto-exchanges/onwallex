FROM  golang:1.23.4-alpine3.19 AS builder
# RUN apk add --no-cache git
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN  go build -o /go/bin/app cmd/main.go

CMD [ "/go/bin/app" ]
#final stage
FROM alpine:latest
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
