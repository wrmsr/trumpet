FROM postgres:9.5.4

RUN apt-get update && \
  apt-get install -y \
    apt-utils \
    vim

RUN \
  sed -i 's/^#wal_level = minimal/wal_level = logical/' /usr/share/postgresql/postgresql.conf.sample && \
  sed -i 's/^#max_replication_slots = 0/max_replication_slots = 10/' /usr/share/postgresql/postgresql.conf.sample && \
  sed -i 's/^#max_wal_senders = 0/max_wal_senders = 10/' /usr/share/postgresql/postgresql.conf.sample && \
  echo 'echo "host replication all 0.0.0.0/0 $authMethod" >> "$PGDATA/pg_hba.conf"' > /docker-entrypoint-initdb.d/add-replication-hba.sh

