FROM postgres:alpine

ENV POSTGRES_PASSWORD=root
ENV POSTGRES_DB=api

RUN apk update \
    apk upgrade 

COPY create_tables.sql /docker-entrypoint-initdb.d/