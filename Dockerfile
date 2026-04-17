FROM mysql:9.6

ENV MYSQL_ROOT_PASSWORD=abcd@1234
ENV MYSQL_DATABASE=user_db

COPY ./users.sql /docker-entrypoint-users.d/