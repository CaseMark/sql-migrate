#!/bin/bash

# Tweak PATH for Travis
export PATH=$PATH:$HOME/gopath/bin

OPTIONS="-config=test-integration/dbconfig.yml -env sqlite"

set -ex

./sql-migrate/sql-migrate status $OPTIONS
./sql-migrate/sql-migrate up $OPTIONS
./sql-migrate/sql-migrate down $OPTIONS
./sql-migrate/sql-migrate redo $OPTIONS
./sql-migrate/sql-migrate status $OPTIONS

# Should have used the custom migrations table
sqlite3 test.db "SELECT COUNT(*) FROM migrations"
