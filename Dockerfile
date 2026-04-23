FROM mysql:9.6



COPY ./database_init.sql /docker-entrypoint-initdb.d/

