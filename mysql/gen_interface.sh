#!/usr/bin/env bash

LIB_ROOT=$(find . -type d -maxdepth 1 -name 'mysql-binary-log-events-*' | head -n 1)
swig -go -cgo -c++ -intgosize 64  -I"$LIB_ROOT/bindings/include" -I"$LIB_ROOT/libbinlogevents/include" mysql_events.i