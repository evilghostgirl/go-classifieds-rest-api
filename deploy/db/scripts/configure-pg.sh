#!/usr/bin/env bash

sed -i "s/^\(listen_addresses .*\)/# Commented out  \1/" /var/lib/postgresql/data/postgresql.conf; 
echo "listen_addresses = '*'" >> /var/lib/postgresql/data/postgresql.conf;
su - postgres;
pg_createcluster 12.2 cluster2;
pg_ctlcluster 12.2 cluster2 start;