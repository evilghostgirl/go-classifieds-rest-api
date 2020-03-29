FROM postgres
RUN sed -i "s/^\(listen_addresses .*\)/# Commented out  \1/"  /var/lib/postgresql/data/postgresql.conf;  
RUN echo "listen_addresses = '*'" >>  /var/lib/postgresql/data/postgresql.conf;
RUN service postgresql restart