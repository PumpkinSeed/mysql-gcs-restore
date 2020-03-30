# mysql-gcs-restore
Restore MySQL database from GCS

### Usage

#### Add GOOGLE_APPLICATION_CREDENTIALS

- Add credentials' json as volume
- Set the environment variable about the location. ex.: `GOOGLE_APPLICATION_CREDENTIALS=/etc/creds.json`

#### Required environment variables

- `BUCKET_NAME=bucket_name`
- `OBJECT_NAME=dump.sql`
- `DB_HOST=127.0.0.1 or mysql.company.com or whatever`
- `DB_USER=user`
- `DB_PASSWORD=password`
- `DB_NAME=db_test`

#### Use Docker image

```
docker run pumpkinseed/mysql-gcs-restore // ADD THE DETAILS FROM ABOVE
```

##### or use docker-compose

```
  restore:
    container_name: restore
    image: mysql-gcs-restore
    environment:
      BUCKET_NAME: bucket_name
      OBJECT_NAME: dump.sql
      DB_HOST: mysql
      DB_USER: user
      DB_PASSWORD: password
      DB_NAME: db_test
      GOOGLE_APPLICATION_CREDENTIALS: '/etc/creds.json'
    volumes:
      - ./creds.json:/etc/creds.json
```