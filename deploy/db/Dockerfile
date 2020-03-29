FROM postgres:12.2
# RUN sed -i "s/^\(listen_addresses .*\)/# Commented out  \1/" /var/lib/postgresql/data/postgresql.conf; \
# echo "listen_addresses = '*'" >> /var/lib/postgresql/data/postgresql.conf; \
RUN su - postgres; \
pg_createcluster 12 cluster2; \
pg_ctlcluster 12 cluster2 start;
RUN service postgresql restart