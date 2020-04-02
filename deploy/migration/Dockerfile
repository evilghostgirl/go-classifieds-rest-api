FROM migrate/migrate
WORKDIR /sql
COPY ./deploy/migration/migrations .
RUN migrate -verbose -source FILE://deploy/migration/migrations -database "postgres://hazelina:secret@db/postgres?sslmode=disable" up