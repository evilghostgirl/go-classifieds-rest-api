FROM migrate/migrate
WORKDIR /sql
COPY ./deploy/migration/migrations .
RUN migrate -verbose -source FILE://deploy/migration/migrations -database "postgres://hazelina:secret@172.27.0.3/postgres?sslmode=disable" up