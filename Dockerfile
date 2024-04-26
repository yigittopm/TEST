FROM golang:1.22 as builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./cmd/wl-auth/main.go

## STEP 2
FROM scratch

WORKDIR /root

COPY .env ./
COPY --from=0 /app/main ./

EXPOSE 8080

CMD [ "./main" ]
