# auth-microservice


# db
## db installation (one-time use)

      sudo -i -u postgres
      psql -U postgres
      postgres=# create database authmcrsrv_db;
      postgres=# create database authmcrsrv_db_test;
      postgres=# create user authmcrsrv_user with encrypted password 'authmcrsrv_pass';
      postgres=# grant all privileges on database authmcrsrv_db_test to authmcrsrv_user;
      postgres=# grant all privileges on database authmcrsrv_db to authmcrsrv_user;
      alter database authmcrsrv_db owner to authmcrsrv_user;
      alter database authmcrsrv_db_test owner to authmcrsrv_user;
      alter schema public owner to authmcrsrv_user;

after that, use this to connect to db in cli

      psql -U authmcrsrv_user -d authmcrsrv_db


## clear db to freshly run migrations

     drop table users;
     drop table goose_db_version;


## migrations


create new migration

      GOOSE_MIGRATION_DIR="internal/storage/migrations" goose create new_table_orders sql

run all migrations

      GOOSE_MIGRATION_DIR="internal/storage/migrations" GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgresql://authmcrsrv_user:authmcrsrv_pass@127.0.0.1:5432/authmcrsrv_db?sslmode=disable" goose up
      GOOSE_MIGRATION_DIR="internal/storage/migrations" GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgresql://authmcrsrv_user:authmcrsrv_pass@127.0.0.1:5432/authmcrsrv_db_test?sslmode=disable" goose up

