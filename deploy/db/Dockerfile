FROM postgres:12.2
# RUN sed -i "s/^\(listen_addresses .*\)/# Commented out  \1/" /var/lib/postgresql/data/postgresql.conf; \
# echo "listen_addresses = '*'" >> /var/lib/postgresql/data/postgresql.conf; \
# RUN curl 
COPY deploy/db/scripts/init.sql /docker-entrypoint-initdb.d/
COPY deploy/db/scripts/check_postgres-2.25.0 /usr/local/bin
RUN apt-get update;
RUN apt-get install curl wget -y;


RUN su - postgres; \
pg_createcluster 12 cluster2; \
pg_ctlcluster 12 cluster2 start;
RUN service postgresql restart
# ./check_postgres.pl --host=localhost --dbname=postgres --dbuser=hazelina --dbpass=secret --action=connection

COPY deploy/db/scripts/register-in-consul.sh /etc/cont-init.d/userconf

# {
#     "Name": "database-postgres",
#     "Address": "database12", 
#     "Check": {
#         "Http":    "http://app:3000/health",
#         "Interval": "10s"
#     }
# }