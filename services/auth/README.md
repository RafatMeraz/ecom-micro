## Migration
In this project, we will use postgres as our main database. For migration, we are using goose.

To start with migration, you need to set up two environment variables;
```
export GOOSE_DRIVER=postgres 
export GOOSE_DBSTRING="user=root password=root@pass host=localhost dbname=shop_management sslmode=disable"
```

To add new migration file
```
goose add name_of_table sql
```

To migrate, run

```
cd db/migrations
goose up
```

To down then migration, use

```
cd db/migrations
goose down
```