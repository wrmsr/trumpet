#!/usr/bin/env bash

INCLUDE_PATH=$(pkg-config mysqlclient --cflags | egrep -o '\-I[^ ]+' | cut -c3-)
wget https://raw.githubusercontent.com/mysql/mysql-server/5.7/include/hash.h > "$INCLUDE_PATH/hash.h"
