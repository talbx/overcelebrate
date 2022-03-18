FROM arm32v7/golang:1.17.8-alpine3.14 as gobuild

WORKDIR /app
RUN apk update
RUN apk upgrade
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM scratch
WORKDIR /app
COPY --from=gobuild /app/app .
COPY --from=gobuild /app/config.yaml .
COPY --from=gobuild /app/bdays.yaml .
COPY --from=gobuild /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./app"]
