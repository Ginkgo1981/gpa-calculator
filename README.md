
### dev
1. create database `gpa_dev`
2. migrate database `make migrate-up`
3. run `make dev` 

### deploy
1. change configuration file in `app/gpa/etc/gpa-api.dev.yaml` 
2. run docker-compose by `make start`

### test
1. create database `create database gpa_test`
2. import sql file `app/gpa/test/gpa_test.sql`
3. change configurations in `app/gpa/test/gpa-api.test.yaml`
4. run test