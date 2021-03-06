#!/usr/bin/env bash

# Fancy macros in pg_type.h break cgo.
# Do I care why? No, life is too short to debug C macros.
# Easier to just run a script to extract type defines and be done with it.

cat `pg_config --includedir-server`/catalog/pg_type.h | \
awk '
BEGIN {
	print "package postgres"
	print "import \"C\""
	print "// WARNING: This file is auto-generated! Do not edit."
	print "const ("
}

/^#define[ \t]+.+OID[ \t]+[0-9]+/ { printf("%s = %s\n", $2, $3) }

END {
	print ")"
}
' > pgtype.go

gofmt -w pgtype.go
