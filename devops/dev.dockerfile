# fast image for fast feedback
FROM alpine:latest

WORKDIR /app

COPY . .

EXPOSE 8081

CMD ["/app/cmd/server/server_linux_amd64"]