FROM alpine

WORKDIR /var/www/

# needed for project dir lookup
ADD ./go.mod /var/www

ADD ./.env /var/www
ADD ./migrations /var/www/

ADD ./cmd/server/server_linux_amd64 /var/www/

CMD ["/var/www/server_linux_amd64"]