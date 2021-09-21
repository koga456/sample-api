FROM --platform=linux/amd64 library/mysql:8.0.25

ENV MYSQL_DATABASE todo

COPY custom.cnf /etc/mysql/conf.d/

COPY sql /docker-entrypoint-initdb.d
