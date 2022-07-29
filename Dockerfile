FROM golang:1.8.3

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www


EXPOSE 4000
ENTRYPOINT ["/var/www"]