FROM arm32v7/golang as gobuild

WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
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
