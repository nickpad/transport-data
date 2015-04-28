#!/bin/bash

set -e

DBNAME="data/nsw.db"

rm -f $DBNAME
cat schema.sql | sqlite3 $DBNAME
csvsql --insert --no-create --db sqlite:///$DBNAME data/calendar.txt
csvsql --insert --no-create --db sqlite:///$DBNAME data/calendar_dates.txt
csvsql --insert --no-create --db sqlite:///$DBNAME data/routes.txt
csvsql --insert --no-create --db sqlite:///$DBNAME data/stop_times.txt
csvsql --insert --no-create --db sqlite:///$DBNAME data/stops.txt
csvsql --insert --no-create --db sqlite:///$DBNAME data/trips.txt
