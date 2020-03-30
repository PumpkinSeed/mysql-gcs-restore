#!/bin/sh

mysql-gcs-restore
/usr/bin/mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_NAME < /tmp/backup.sql