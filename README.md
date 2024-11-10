# auth-microservice


# db
## 
must have these env vars

    POSTGRES_DB 
    POSTGRES_USER
    POSTGRES_PASSWORD

after that, use this to connect to db in cli

      psql -U authmcrsrv_user -d authmcrsrv_db


## clear db to freshly run migrations

     drop table users;
     drop table goose_db_version;


## migrations


create new migration

      GOOSE_MIGRATION_DIR="migrations" goose create new_table_orders sql

run all migrations

      GOOSE_MIGRATION_DIR="migrations" GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgresql://authmcrsrv_user:authmcrsrv_pass@127.0.0.1:5432/authmcrsrv_db?sslmode=disable" goose up
      GOOSE_MIGRATION_DIR="migrations" GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgresql://authmcrsrv_user:authmcrsrv_pass@127.0.0.1:5432/authmcrsrv_db_test?sslmode=disable" goose up

