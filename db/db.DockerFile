FROM postgres

COPY ./docs/init.sql /docker-entrypoint-initdb.d/10-init.sql
CMD ["docker-entrypoint.sh", "postgres"]